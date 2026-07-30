package repository

import (
	"context"
	"strings"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type CatalogItemRepository interface {
	GetCatalogItems(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.CatalogItem, int64, error)
	GetCatalogItemByID(ctx context.Context, tenantID string, id string) (*model.CatalogItem, error)
	CreateCatalogItem(ctx context.Context, ci *model.CatalogItem) error
	UpdateCatalogItem(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	DeleteCatalogItem(ctx context.Context, tenantID string, id string) error
	// method untuk dashboard, hanya menampilkan jumlah category item
	CountCatalogItems(ctx context.Context, tenantID string) (int64, error)
}

// sturct implementasi
type catalogItemRepository struct {
	db *gorm.DB
}

// constructor
func NewCatalogItemRepository(db *gorm.DB) CatalogItemRepository {
	return &catalogItemRepository{
		db: db,
	}
}

// struct method
func (r *catalogItemRepository) GetCatalogItems(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.CatalogItem, int64, error) {
	var ci []model.CatalogItem
	var total int64

	// query default
	query := r.db.WithContext(ctx).
		Model(model.CatalogItem{}).
		Joins("LEFT JOIN catalog_categories ON catalog_categories.id = catalog_items.category_id").
		Preload("Tenant").
		Preload("CatalogCategory").
		Where("catalog_items.tenant_id = ?", tenantID)

	// cek name kosong atau tidak
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
			catalog_items.name LIKE ?
			OR catalog_items.description LIKE ?
			OR catalog_categories.name LIKE ?
		)
		`,
			like,
			like,
			like,
		)
	}

	// hitung jumlah data sebelum pagination
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Whitelist sort field
	allowedSort := map[string]string{
		"name":          "catalog_items.name",
		"category_name": "catalog_categories.name",
		"description":   "catalog_items.description",
	}

	// Default sort
	sortField := "catalog_items.name"

	// Jika field valid, gunakan field hasil mapping
	if value, ok := allowedSort[req.Sort]; ok {
		sortField = value
	}

	// Validasi order
	order := "DESC"
	if strings.ToUpper(req.Order) == "ASC" {
		order = "ASC"
	}

	// Sorting dengan whitelist
	query = query.Order(sortField + " " + order)

	// Pagination
	offset := (req.Page - 1) * req.Limit

	err = query.
		Offset(offset).
		Limit(req.Limit).
		Find(&ci).Error

	if err != nil {
		return nil, 0, err
	}

	return ci, total, nil
}

func (r *catalogItemRepository) GetCatalogItemByID(ctx context.Context, tenantID string, id string) (*model.CatalogItem, error) {
	var ci model.CatalogItem
	err := r.db.WithContext(ctx).Preload("Tenant").Preload("CatalogCategory").Where("tenant_id = ?", tenantID).First(&ci, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &ci, nil
}

func (r *catalogItemRepository) CreateCatalogItem(ctx context.Context, ci *model.CatalogItem) error {
	return r.db.WithContext(ctx).Create(ci).Error
}

func (r *catalogItemRepository) UpdateCatalogItem(ctx context.Context, tenantID string, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.CatalogItem{}).Where("id = ? AND tenant_id = ?", id, tenantID).Updates(updateMap).Error
}

func (r *catalogItemRepository) DeleteCatalogItem(ctx context.Context, tenantID string, id string) error {
	return r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&model.CatalogItem{}).Error
}

func (r *catalogItemRepository) CountCatalogItems(ctx context.Context, tenantID string) (int64, error) {
	var total int64
	err := r.db.
		WithContext(ctx).
		Model(model.CatalogItem{}).
		Where("tenant_id = ?", tenantID).
		Count(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}
