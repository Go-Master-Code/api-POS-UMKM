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

// interface
type ExpenseService interface {
	GetAllExpenses(ctx context.Context, query dto.GetAllExpenseQuery) ([]dto.ExpensesResponse, int64, error)
	GetExpenseByID(ctx context.Context, id string) (dto.ExpensesResponse, error)
	CreateExpense(ctx context.Context, req dto.CreateExpenseRequest) (dto.ExpensesResponse, error)
}

// struct implementasi
type expenseService struct {
	expenseRepo     repository.ExpenseRepository
	expenseItemRepo repository.ExpenseItemRepository
	db              *gorm.DB
	// log
	activityLogService ActivityLogService // jangan pakai package service, karena kedua file ini ada di dalam package yang sama (service)
}

// constructor
func NewExpenseServie(expenseRepo repository.ExpenseRepository, expenseItemRepo repository.ExpenseItemRepository, db *gorm.DB, activityLogService ActivityLogService) ExpenseService {
	return &expenseService{
		expenseRepo:        expenseRepo,
		expenseItemRepo:    expenseItemRepo,
		db:                 db,
		activityLogService: activityLogService,
	}
}

// struct method
func (s *expenseService) GetAllExpenses(ctx context.Context, query dto.GetAllExpenseQuery) ([]dto.ExpensesResponse, int64, error) {
	// ambil tenantID dari jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	expenses, total, err := s.expenseRepo.GetAllExpenses(ctx, tenantID, query)
	if err != nil {
		return nil, 0, err
	}

	// convert model to dto
	expensesDTO := helper.ConvertToDTOExpensePlural(expenses)

	return expensesDTO, total, nil
}

func (s *expenseService) GetExpenseByID(ctx context.Context, id string) (dto.ExpensesResponse, error) {
	// get tenantID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	expense, err := s.expenseRepo.GetExpenseByID(ctx, tenantID, id)
	if err != nil {
		return dto.ExpensesResponse{}, err
	}

	// convert model to dto
	expenseDTO := helper.ConvertToDTOExpenseSingle(expense)
	return expenseDTO, nil
}

func (s *expenseService) CreateExpense(ctx context.Context, req dto.CreateExpenseRequest) (dto.ExpensesResponse, error) {
	// inisiasi tx
	tx := s.db.Begin()

	if tx.Error != nil { // cek apakah gagal begin transaction
		return dto.ExpensesResponse{}, tx.Error
	}

	// safety rollback jika panic
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// ambil data tenant dan user dari jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	userID := ctx.Value(constants.ContextUserID).(string)

	// ========================================
	// GENERATE EXPENSE NUMBER
	// ========================================
	expenseNumber := fmt.Sprintf(
		"EXP-%d",
		time.Now().Unix(),
	)

	// ========================================
	// CREATE EXPENSE HEADER -> master expense
	// ========================================
	expense := model.Expenses{
		ID:            uuid.NewString(),
		TenantID:      tenantID,
		PaymentMethod: req.PaymentMethod,
		ExpenseNumber: expenseNumber,
		Notes:         req.Notes,
		CreatedBy:     userID,
	}

	// simpan data master expense
	err := s.expenseRepo.CreateExpense(ctx, tx, &expense)
	if err != nil {
		tx.Rollback() // jika terjadi error saat input master expense, rollback
		return dto.ExpensesResponse{}, err
	}

	// ========================================
	// LOOP EXPENSE ITEMS
	// ========================================
	for _, item := range req.Items {
		// validasi qty
		if item.Qty < 1 { // qty minimal 1
			tx.Rollback()
			return dto.ExpensesResponse{}, errors.New(
				"item qty cannot be less than 1",
			)
		}

		// validasi harga
		if item.UnitPrice < 1 { // harga minimal 1
			tx.Rollback()
			return dto.ExpensesResponse{}, errors.New(
				"unit price cannot be less than 1",
			)
		}

		// hitung subtotal item expense
		subtotal := item.UnitPrice * item.Qty

		// create expense item
		expenseItem := model.ExpenseItem{
			ID:                uuid.NewString(),
			ExpenseID:         expense.ID,
			ExpenseCategoryID: item.ExpenseCategoryID,
			Description:       item.Description,
			Qty:               item.Qty,
			Unit:              item.Unit,
			UnitPrice:         item.UnitPrice,
			Subtotal:          subtotal,
		}

		// insert data ke database
		err := s.expenseItemRepo.CreateExpenseItem(ctx, tx, &expenseItem)

		if err != nil {
			tx.Rollback()
			return dto.ExpensesResponse{}, err
		}

		// akumulasi nilai total amount untuk master expense
		expense.TotalAmount += expenseItem.Subtotal
	}

	// update total amount ke master expense
	err = s.expenseRepo.UpdateTotalAmount(ctx, tx, expense.ID, expense.TotalAmount)
	if err != nil {
		tx.Rollback()
		return dto.ExpensesResponse{}, err
	}

	// COMMIT TRANSACTION
	err = tx.Commit().Error

	if err != nil {
		return dto.ExpensesResponse{}, err
	}

	// ========================================
	// ACTIVITY LOG
	// ========================================
	//
	// Transaction sudah COMMIT.
	// Kegagalan activity log tidak boleh membuat
	// frontend menganggap transaksi gagal.
	err = s.activityLogService.CreateActivityLog( // ignore error
		ctx,
		"EXPENSE",
		"CREATE",
		fmt.Sprintf("Create expense %s", expense.ID),
		expense.ID,            // id uuid
		expense.ExpenseNumber, // yang mudah dipahami manusia misalnya EXP-20260827
	)

	// error log activity jgn kirim error ke client
	if err != nil {
		log.Println("Error log: ", err)
	}

	// get data expense by id untuk preload semua relasi
	newExpense, err := s.expenseRepo.GetExpenseByID(ctx, tenantID, expense.ID)
	if err != nil {
		return dto.ExpensesResponse{}, err
	}

	// ========================================
	// RESPONSE DTO
	// ========================================
	expenseDTO := helper.ConvertToDTOExpenseSingle(newExpense)

	return expenseDTO, nil
}
