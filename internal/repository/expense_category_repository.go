package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ExpenseCategoryRepository interface {
	GetAllExpenseCategoryPerTenant(ctx context.Context, tenantID string) ([]model.ExpenseCategory, error) // tampilkan semua data di combobox
	GetExpenseCategory(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.ExpenseCategory, int64, error)
	// GetExpenseCategoryByID(ctx context.Context, tenantID string, id string) (*model.ExpenseCategory, error)
	// CreateExpenseCategory(ctx context.Context, ec *model.ExpenseCategory) error
	// UpdateExpenseCategory(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	// DeleteExpenseCategory(ctx context.Context, tenantID string, id string) error
}

// struct
type expenseCategoryRepository struct {
	db *gorm.DB
}

// constructor
func NewExpenseCategoryRepository(db *gorm.DB) ExpenseCategoryRepository {
	return &expenseCategoryRepository{
		db: db,
	}
}

// struct method
func (r *expenseCategoryRepository) GetAllExpenseCategoryPerTenant(ctx context.Context, tenantID string) ([]model.ExpenseCategory, error) {
	var expenseCategories []model.ExpenseCategory
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&expenseCategories).Error
	if err != nil {
		return nil, err
	}

	return expenseCategories, nil
}

func (r *expenseCategoryRepository) GetExpenseCategory(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.ExpenseCategory, int64, error) {
	var ec []model.ExpenseCategory
	var total int64

	// query default
	query := r.db.WithContext(ctx).
		Model(&model.ExpenseCategory{}).
		Preload("Tenant").
		Where("tenant_id = ?", tenantID)

	// cek name / deskripsi kosong atau tidak
	if req.Search != "" {
		like := "%" + req.Search + "%"

		// perhatikan baik-baik tanda () di klausa WHERE di bawah, harus sama persis!
		/*
			Apakah tenant = Tenant A?
			YES
			│
			├── Apakah nama/deskripsi/kategori mengandung "Snack"?
			│       YES → tampil
			│       NO  → sembunyikan
			NO
			└── langsung buang
		*/
		query = query.Where(`
		(
			name LIKE ?
			OR description LIKE ?
		)
		`,
			like,
			like,
		)
	}

	// hitung jumlah row
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Whitelist sort field
	allowedSort := map[string]string{
		"name":        "name",
		"description": "description",
	}

	// Default sort
	sortField := "name"

	// Jika field valid, gunakan field hasil mapping
	if value, ok := allowedSort[req.Sort]; ok {
		sortField = value
	}

	// order asc saja, berdasarkan abjad field "name" from A to Z
	order := "ASC"

	// Sorting dengan whitelist
	query = query.Order(sortField + " " + order)

	// Pagination
	offset := (req.Page - 1) * req.Limit

	err = query.
		Offset(offset).
		// Limit(req.Limit). untuk sekarang jangan pakai limit karena jadinya tidak akan muncul semua expense category nya pada combobox
		Find(&ec).Error

	if err != nil {
		return nil, 0, err
	}

	return ec, total, nil
}
