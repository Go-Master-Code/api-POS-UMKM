package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ExpenseRepository interface {
	GetAllExpenses(ctx context.Context, tenantID string, query dto.GetAllExpenseQuery) ([]model.Expenses, int64, error) // per tenant
	GetExpenseByID(ctx context.Context, tenantID string, id string) (*model.Expenses, error)
	CreateExpense(ctx context.Context, tx *gorm.DB, expense *model.Expenses) error
	UpdateTotalAmount(ctx context.Context, tx *gorm.DB, id string, totalAmount float64) error
}

// struct implementasi
type expenseRepository struct {
	db *gorm.DB
}

// constructor
func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{
		db: db,
	}
}

// struct method
func (r *expenseRepository) GetAllExpenses(ctx context.Context, tenantID string, query dto.GetAllExpenseQuery) ([]model.Expenses, int64, error) {
	var expenses []model.Expenses
	var total int64

	// pagination sudah dilakukan di handler
	offset := (query.Page - 1) * query.Limit

	// base query
	baseQuery := r.db.WithContext(ctx).Model(&model.Expenses{}).Where("tenant_id = ?", tenantID)

	// filter payment status
	if query.PaymentMethod != "" {
		baseQuery = baseQuery.Where("payment_method = ?", query.PaymentMethod)
	}

	// search amount
	if query.Search != "" {
		baseQuery = baseQuery.Where("total_amount = ?", query.Search)
	}

	// count total row
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
		"total_amount":   "total_amount",
		"payment_method": "payment_method",
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
		Preload("User").
		Preload("ExpenseItems").
		Preload("ExpenseItems.ExpenseCategory").
		Order(sortField + " " + sortOrder).
		Limit(query.Limit).Offset(offset).Find(&expenses).Error

	if err != nil {
		return nil, 0, err
	}

	// jika semua query sukses
	return expenses, total, nil
}

func (r *expenseRepository) GetExpenseByID(ctx context.Context, tenantID string, id string) (*model.Expenses, error) {
	var expense model.Expenses
	err := r.db.WithContext(ctx).
		Preload("Tenant").
		Preload("User").
		Preload("ExpenseItems").
		Preload("ExpenseItems.ExpenseCategory").
		Where("tenant_id = ? and id = ?", tenantID, id).
		First(&expense).Error
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

// CreateExpense pakai tx bukan r.db karena expense dan expense item harus dalam 1 transaction yang sama
func (r *expenseRepository) CreateExpense(ctx context.Context, tx *gorm.DB, expense *model.Expenses) error {
	return tx.WithContext(ctx).Create(expense).Error
}

func (r *expenseRepository) UpdateTotalAmount(ctx context.Context, tx *gorm.DB, id string, totalAmount float64) error {
	result := tx.WithContext(ctx).
		Model(&model.Expenses{}).
		Where("id = ?", id).
		Update("total_amount", totalAmount)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 { // jika tidak ada row yang terupdate valuenya
		return gorm.ErrRecordNotFound
	}

	return nil
}
