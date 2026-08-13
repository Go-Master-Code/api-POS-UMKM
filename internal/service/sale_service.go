package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
	===TARGET FLOW===
	BEGIN TRANSACTION
	↓
	create sale
	↓
	loop items
		↓
		get variant
		↓
		validate stock
		↓
		create sale item
		↓
		create stock movement
	↓
	update grand total
	↓
	COMMIT
*/

// interface
type SaleService interface {
	GetAllSales(ctx context.Context, query dto.GetAllSalesQuery) ([]dto.SaleResponse, int64, error)
	CreateSale(ctx context.Context, req dto.CreateSaleRequest) (dto.SaleResponse, error)
	GetSaleByID(ctx context.Context, id string) (dto.SaleResponse, error)
}

// struct implementasi
type saleService struct {
	db                *gorm.DB // ada db di service karena: transaction begin/commit/rollback dilakukan di layer service
	saleRepo          repository.SaleRepository
	saleItemRepo      repository.SaleItemRepository
	itemVariantRepo   repository.ItemVariantRepository
	stockMovementRepo repository.StockMovementRepository

	// log
	activityLogService ActivityLogService // jangan pakai package service, karena kedua file ini ada di dalam package yang sama (service)
}

// constructor -> ada db karena untuk transaction
func NewSaleService(
	db *gorm.DB,
	saleRepo repository.SaleRepository,
	saleItemRepo repository.SaleItemRepository,
	itemVariantRepo repository.ItemVariantRepository,
	stockMovementRepo repository.StockMovementRepository,
	activityLogService ActivityLogService,
) SaleService {
	return &saleService{
		db:                 db,
		saleRepo:           saleRepo,
		saleItemRepo:       saleItemRepo,
		itemVariantRepo:    itemVariantRepo,
		stockMovementRepo:  stockMovementRepo,
		activityLogService: activityLogService,
	}
}

// stuct method
func (s *saleService) GetAllSales(ctx context.Context, query dto.GetAllSalesQuery) ([]dto.SaleResponse, int64, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data from repository
	sales, total, err := s.saleRepo.GetAllSales(ctx, tenantID, query)

	if err != nil {
		return nil, 0, err
	}

	// convert model to dto
	salesDTO := helper.ConvertToDTOSalePlural(sales)

	// jika semua sukses
	return salesDTO, total, nil
}

func (s *saleService) CreateSale(ctx context.Context, req dto.CreateSaleRequest) (dto.SaleResponse, error) {
	// ========================================
	// BEGIN DATABASE TRANSACTION
	// ========================================

	log.Println("SALE DEBUG 01 - BEGIN")
	tx := s.db.Begin() // begin transaction

	if tx.Error != nil { // cek apakah gagal begin transaction
		return dto.SaleResponse{}, tx.Error
	}
	log.Println("SALE DEBUG 02 - BEGIN SUCCESS")

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// ========================================
	// AMBIL DATA USER DARI CONTEXT JWT
	// ========================================
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	userID := ctx.Value(constants.ContextUserID).(string)

	// ========================================
	// GENERATE INVOICE NUMBER
	// ========================================
	invoiceNumber := fmt.Sprintf(
		"INV-%d",
		time.Now().Unix(),
	)

	// ========================================
	// CREATE SALE HEADER -> master sale
	// ========================================
	sale := model.Sale{
		ID:             uuid.NewString(),
		TenantID:       tenantID,
		InvoiceNumber:  invoiceNumber,
		CustomerName:   req.CustomerName,
		CashierID:      userID,
		DiscountAmount: req.DiscountAmount,
		PaymentMethod:  req.PaymentMethod,
		AmountReceived: req.AmountReceived,
		// PaymentStatus:  req.PaymentStatus, ditentukan nanti setelah grandTotal selesai dihitung
		Notes: req.Notes,
	}

	// simpan sale header / data master
	err := s.saleRepo.CreateSale(ctx, tx, &sale)
	if err != nil {
		tx.Rollback() // jika terjadi error saat input data header / master sale, rollback
		return dto.SaleResponse{}, err
	}

	// ========================================
	// LOOP SALE ITEMS
	// ========================================
	for _, item := range req.Items { // iterasi ke sale items -> lihat format dtoCreateSaleRequest
		// ----------------------------------------
		// AMBIL ITEM VARIANT DARI DATABASE
		// ----------------------------------------
		// Harga dan informasi item harus trusted
		// dari database, bukan dari frontend.
		variant, err := s.itemVariantRepo.GetItemVariantByIDForUpdate(ctx, tx, tenantID, item.ItemVariantID)
		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}
		log.Println("SALE DEBUG 03 - VARIANT LOCKED:", variant.ID)

		// ----------------------------------------
		// VALIDASI QTY
		// ----------------------------------------
		if item.Qty <= 0 {
			tx.Rollback()
			return dto.SaleResponse{}, errors.New(
				"quantity must be greater than zero",
			)
		}

		// ========================================
		// VALIDASI STOK
		// ========================================

		currentStock, err := s.stockMovementRepo.GetCurrentStockForUpdate( // wajib pakai param tx dan jalankan di repo pakai tx, bukan r.db
			ctx,
			tenantID,
			tx,
			variant.ID,
		)
		log.Println("SALE DEBUG 04 - CURRENT STOCK:", currentStock)

		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}

		// stok tidak boleh minus
		if currentStock < item.Qty {
			tx.Rollback()
			return dto.SaleResponse{},
				errors.New("insufficient stock")
		}

		// ----------------------------------------
		// VALIDASI DISCOUNT ITEM
		// ----------------------------------------

		grossSubtotal := item.Qty * variant.SellingPrice

		if item.DiscountAmount < 0 { // diskon negatif
			tx.Rollback()
			return dto.SaleResponse{}, errors.New(
				"item discount cannot be negative",
			)
		}

		if item.DiscountAmount > grossSubtotal { // diskon tidak boleh > subtotal
			tx.Rollback()
			return dto.SaleResponse{}, errors.New(
				"item discount cannot exceed item subtotal",
			)
		}

		// ========================================
		// HITUNG SUBTOTAL ITEM
		// ========================================

		subtotal := grossSubtotal - item.DiscountAmount

		// ========================================
		// CREATE SALE ITEM
		// gunakan snapshot agar histori immutable
		// ========================================

		saleItem := model.SaleItem{
			ID:                  uuid.NewString(),
			TenantID:            tenantID,
			SaleID:              sale.ID,
			ItemVariantID:       variant.ID,
			ItemNameSnapshot:    variant.Item.Name,
			VariantNameSnapshot: variant.VariantName,
			SKUSnapshot:         variant.SKU,
			Qty:                 item.Qty,
			UnitPrice:           variant.SellingPrice,
			// DiscountAmount:      item.DiscountAmount,
			DiscountAmount: 0, // sementara dibuat 0 diskon untuk setiap sale item karena diskon hanya ada di master sales
			Subtotal:       subtotal,
		}

		// simpan sale item
		err = s.saleItemRepo.CreateSaleItem(ctx, tx, &saleItem)

		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}

		log.Println("SALE DEBUG 05 - SALE ITEM CREATED")

		// ========================================
		// CREATE STOCK MOVEMENT
		// stok keluar = qty negatif
		// ========================================

		movement := model.StockMovement{
			ID:            uuid.NewString(),
			TenantID:      tenantID,
			ItemVariantID: variant.ID,
			MovementType:  constants.MovementSale,
			Qty:           -item.Qty,
			ReferenceType: "SALE",
			ReferenceID:   sale.ID,
			Notes:         "sale transaction",
			CreatedBy:     userID,
		}

		log.Println("SALE DEBUG 06A - BEFORE CREATE MOVEMENT")
		// simpan stock movement
		err = s.stockMovementRepo.CreateMovement(ctx, tx, &movement)

		if err != nil {
			tx.Rollback()
			return dto.SaleResponse{}, err
		}
		log.Println("SALE DEBUG 06B - AFTER CREATE MOVEMENT")

		// ========================================
		// TAMBAHKAN SUBTOTAL ITEM KE SALE
		// ========================================
		// increment sale.subtotal
		sale.Subtotal += subtotal
	}

	// ========================================
	// VALIDASI SALE DISCOUNT
	// ========================================

	if sale.DiscountAmount < 0 {
		tx.Rollback()
		return dto.SaleResponse{}, errors.New(
			"sale discount cannot be negative",
		)
	}

	if sale.DiscountAmount > sale.Subtotal {
		tx.Rollback()
		return dto.SaleResponse{}, errors.New(
			"sale discount cannot exceed subtotal",
		)
	}

	// ========================================
	// HITUNG TAX
	// ========================================
	//
	// Untuk sementara masih 10%.
	// Nanti kita pindahkan ke tenant/settings
	// jika memang dibutuhkan.

	taxableAmount := sale.Subtotal - sale.DiscountAmount
	sale.TaxAmount = taxableAmount / 10 // skenario tax=10%

	// ========================================
	// HITUNG FINAL GRAND TOTAL
	// ========================================

	// update grand total setelah diskon ditambah pajak
	sale.GrandTotal = taxableAmount + sale.TaxAmount

	// ========================================
	// VALIDASI PAYMENT
	// ========================================
	switch sale.PaymentMethod {
	case constants.PaymentMethodCash:
		if sale.PaymentMethod == constants.PaymentMethodCash {
			if req.AmountReceived < sale.GrandTotal {
				tx.Rollback()

				return dto.SaleResponse{}, errors.New("amount received is less than grand total")
			}
		}

		// pembayaran cash berhasil
		sale.PaymentStatus = constants.PaymentStatusPaid

	case constants.PaymentMethodQRIS,
		constants.PaymentMethodTransfer,
		constants.PaymentMethodDebit,
		constants.PaymentMethodKredit:

		// Untuk MVP, pembayaran dianggap sudah dikonfirmasi di kasir
		sale.AmountReceived = sale.GrandTotal
		sale.PaymentStatus = constants.PaymentStatusPaid

	default:
		tx.Rollback()

		return dto.SaleResponse{}, errors.New("invalid payment method")
	}

	// validasi nominal bayar jika metode pembayaran CASH

	err = tx.
		WithContext(ctx).
		Model(&sale).
		Updates(map[string]any{
			"subtotal":        sale.Subtotal,
			"discount_amount": sale.DiscountAmount,
			"tax_amount":      sale.TaxAmount,
			"grand_total":     sale.GrandTotal,
			"payment_status":  sale.PaymentStatus,
			"amount_received": sale.AmountReceived,
		}).Error

	if err != nil {
		tx.Rollback()
		return dto.SaleResponse{}, err
	}

	// ========================================
	// COMMIT TRANSACTION
	// ========================================

	log.Println("SALE DEBUG 07 - BEFORE COMMIT")
	err = tx.Commit().Error

	if err != nil {
		return dto.SaleResponse{}, err
	}
	log.Println("SALE DEBUG 08 - COMMIT SUCCESS")

	// ========================================
	// ACTIVITY LOG
	// ========================================
	//
	// Transaction sudah COMMIT.
	// Kegagalan activity log tidak boleh membuat
	// frontend menganggap transaksi gagal.
	err = s.activityLogService.CreateActivityLog( // ignore error
		ctx,
		"SALES",
		"CREATE",
		fmt.Sprintf("Create Sales %s", sale.InvoiceNumber),
		sale.ID,            // id uuid
		sale.InvoiceNumber, // yang mudah dipahami manusia misalnya P-RETUR-1781140525
	)

	// error log activity jgn kirim error ke client
	if err != nil {
		log.Println("Error log: ", err)
	}

	// get data sale by id untuk preload semua relasi
	newSale, err := s.saleRepo.GetSaleByID(ctx, tenantID, sale.ID)
	if err != nil {
		return dto.SaleResponse{}, err
	}
	log.Println("SALE DEBUG 09 - AFTER COMMIT")

	// ========================================
	// RESPONSE DTO
	// ========================================
	saleDTO := helper.ConvertToDTOSaleSingle(newSale)

	return saleDTO, nil
}

func (s *saleService) GetSaleByID(ctx context.Context, id string) (dto.SaleResponse, error) {
	// ambil tenantID dari context
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	// akses repo
	sale, err := s.saleRepo.GetSaleByID(ctx, tenantID, id)
	if err != nil {
		return dto.SaleResponse{}, err
	}

	// convert model to dto
	saleDTO := helper.ConvertToDTOSaleSingle(sale)
	return saleDTO, nil
}
