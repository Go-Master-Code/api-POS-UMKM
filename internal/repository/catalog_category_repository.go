package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type CatalogCategoryRepository interface {
	GetCatalogCategories(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.CatalogCategory, int64, error)
	GetCatalogCategoryByID(ctx context.Context, tenantID string, id string) (*model.CatalogCategory, error)
	CreateCatalogCategory(ctx context.Context, cc *model.CatalogCategory) error
	UpdateCatalogCategory(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	DeleteCatalogCategory(ctx context.Context, tenantID string, id string) error
}

// struct implementasi
type catalogCategoryRepository struct {
	db *gorm.DB
}

// constructor
func NewCatalogCategoryRepository(db *gorm.DB) CatalogCategoryRepository {
	return &catalogCategoryRepository{
		db: db,
	}
}

// struct method
func (r *catalogCategoryRepository) GetCatalogCategories(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.CatalogCategory, int64, error) {
	var cc []model.CatalogCategory
	var total int64 // untuk return value total data

	// query utama (HARUS MENCANTUMKAN Model())
	query := r.db.WithContext(ctx).Model(&model.CatalogCategory{}).Preload("Tenant").Where("tenant_id = ?", tenantID)

	// jika name tidak kosong
	if req.Search != "" {
		query = query.Where("name LIKE ? OR DATE_FORMAT(created_at, '%d %b %Y') LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	// hitung jumlah data sebelum pagination
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Sorting (sementara, nanti kita whitelist)
	query = query.Order(req.Sort + " " + req.Order)

	// Pagination
	offset := (req.Page - 1) * req.Limit

	err = query.
		Offset(offset).
		Limit(req.Limit).
		Find(&cc).Error

	if err != nil {
		return nil, 0, err
	}

	return cc, total, nil
}

func (r *catalogCategoryRepository) GetCatalogCategoryByID(ctx context.Context, tenantID string, id string) (*model.CatalogCategory, error) {
	var cc model.CatalogCategory
	err := r.db.WithContext(ctx).Preload("Tenant").Where("tenant_id = ?", tenantID).First(&cc, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &cc, nil
}

func (r *catalogCategoryRepository) CreateCatalogCategory(ctx context.Context, cc *model.CatalogCategory) error {
	return r.db.WithContext(ctx).Create(cc).Error
}

func (r *catalogCategoryRepository) UpdateCatalogCategory(ctx context.Context, tenantID string, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.CatalogCategory{}).Where("id = ? AND tenant_id = ?", id, tenantID).Updates(updateMap).Error
}

func (r *catalogCategoryRepository) DeleteCatalogCategory(ctx context.Context, tenantID string, id string) error {
	return r.db.WithContext(ctx).Where("id = ? and tenant_id = ?", id, tenantID).Delete(&model.CatalogCategory{}).Error
}
