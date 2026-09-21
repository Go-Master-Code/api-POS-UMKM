package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ExpenseItemRepository interface {
	CreateExpenseItem(ctx context.Context, tx *gorm.DB, expenseItem *model.ExpenseItem) error
}

// struct implementasi
type expenseItemRepository struct {
	db *gorm.DB
}

// constructor
func NewExpenseItemRepository(db *gorm.DB) ExpenseItemRepository {
	return &expenseItemRepository{
		db: db,
	}
}

// struct method
func (r *expenseItemRepository) CreateExpenseItem(ctx context.Context, tx *gorm.DB, expenseItem *model.ExpenseItem) error {
	return tx.WithContext(ctx).Create(expenseItem).Error
}
