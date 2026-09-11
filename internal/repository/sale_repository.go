package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

/*
	Gambaran besar proses create sales, sale items, dan stock movement
	tx := db.Begin()
	saleRepo.CreateSale(tx)
	saleItemRepo.CreateSaleItem(tx)
	stockRepo.CreateMovement(tx)
	tx.Commit()
	Penjelasan:
		-atomic
		-konsisten
		-aman rollback
		-tidak corrupt
*/

// interface
type SaleRepository interface {
	GetAllSales(ctx context.Context, tenantID string, query dto.GetAllSalesQuery) ([]model.Sale, int64, error)
	CreateSale(ctx context.Context, tx *gorm.DB, sale *model.Sale) error
	GetSaleByID(ctx context.Context, tenantID string, id string) (*model.Sale, error) // perlu tenant isolation agar tenant A tidak bisa akses invoice tenant B
	PaySale(ctx context.Context, tenantID string, saleID string, amountReceived float64) error
}

// struct implementasi
type saleRepository struct {
	db *gorm.DB
}

// constructor
func NewSaleRepository(db *gorm.DB) SaleRepository {
	return &saleRepository{
		db: db,
	}
}

// struct method.

func (r *saleRepository) GetAllSales(ctx context.Context, tenantID string, query dto.GetAllSalesQuery) ([]model.Sale, int64, error) {
	var sales []model.Sale
	var total int64

	// pagination sudah dilakukan di handler

	offset := (query.Page - 1) * query.Limit

	// base query
	baseQuery := r.db.WithContext(ctx).Model(&model.Sale{}).Where("tenant_id = ?", tenantID)

	// filter payment status
	if query.PaymentStatus != "" {
		baseQuery = baseQuery.Where("payment_status = ?", query.PaymentStatus)
	}

	// search
	if query.Search != "" {
		search := "%" + query.Search + "%"
		baseQuery = baseQuery.Where("invoice_number LIKE ? OR customer_name LIKE ?", search, search) // query untuk cari data yang invoice atau customer name nya like ...
	}

	// count total row(s) found -> diperlukan untuk frontend
	err := baseQuery.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	/*
		|--------------------------------------------------------------------------
		| Sorting Whitelist
		|--------------------------------------------------------------------------
		|
		| Hanya field yang terdaftar di sini yang boleh digunakan
		| untuk sorting.
		|
	*/

	allowedSortFields := map[string]string{
		"invoice_number": "invoice_number",
		"customer_name":  "customer_name",
		"grand_total":    "grand_total",
		"created_at":     "created_at",
	}

	// default sorting
	sortField := "created_at"
	sortOrder := "DESC"

	// validasi sort field menggunakan whitelist
	if field, ok := allowedSortFields[query.Sort]; ok {
		sortField = field
	}

	// validasi sort order
	switch query.Order {
	case "asc":
		sortOrder = "ASC"
	case "desc":
		sortOrder = "DESC"
	}

	// get data
	err = baseQuery.Preload("Tenant").
		Preload("Cashier").
		Preload("Customer").
		Preload("SaleItems").
		Preload("SaleItems.Tenant").
		Preload("SaleItems.ItemVariant").
		Order(sortField + " " + sortOrder).
		Limit(query.Limit).Offset(offset).Find(&sales).Error
	if err != nil {
		return nil, 0, err
	}

	// jika semua sukses
	return sales, total, nil
}

// CreateSale pakai tx bukan r.db karena sale, sale item, dan stock movement harus dalam 1 transaction yang sama
func (r *saleRepository) CreateSale(ctx context.Context, tx *gorm.DB, sale *model.Sale) error {
	return tx.WithContext(ctx).Create(sale).Error
}

func (r *saleRepository) GetSaleByID(ctx context.Context, tenantID string, id string) (*model.Sale, error) {
	var sale model.Sale
	err := r.db.
		WithContext(ctx).
		Preload("Tenant").
		Preload("Customer").
		Preload("Cashier").
		Preload("SaleItems").
		Preload("SaleItems.Tenant").      // preload nested relation dari sale item
		Preload("SaleItems.Sale").        // preload nested relation dari sale item
		Preload("SaleItems.ItemVariant"). // preload nested relation dari sale item
		Where("id = ? AND tenant_id = ?", id, tenantID).
		First(&sale).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *saleRepository) PaySale(ctx context.Context, tenantID string, saleID string, amountReceived float64) error {
	var updateMap = map[string]any{}

	// masukkan value updateMap
	updateMap["amount_received"] = amountReceived
	updateMap["payment_status"] = "PAID"

	return r.db.WithContext(ctx).Model(model.Sale{}).Where("id = ? AND tenant_id = ?", saleID, tenantID).Updates(updateMap).Error
}
