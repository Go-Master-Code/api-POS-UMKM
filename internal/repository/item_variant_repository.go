package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ItemVariantRepository interface {
	GetItemVariants(ctx context.Context, tenantID string, catalogItemID string, req dto.PaginationRequest) ([]model.ItemVariant, int64, error)
	CountItemVariants(ctx context.Context, tenantID string) (int64, error) // untuk summary dashboard
	GetItemVariantByID(ctx context.Context, tenantID string, id string) (*model.ItemVariant, error)
	CreateItemVariant(ctx context.Context, iv *model.ItemVariant) error
	UpdateItemVariant(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	DeleteItemVariant(ctx context.Context, tenantID string, id string) error
	// ambil item yang <= low stock
	GetLowStockItems(ctx context.Context, tenantID string) ([]dto.LowStockResponse, error)
	// // menghitung jumlah item low stock
	GetLowStockCount(ctx context.Context, tenantID string) (int64, error)
}

// struct implementasi
type itemVariantRepository struct {
	db *gorm.DB
}

// constructor
func NewItemVariantRepository(db *gorm.DB) ItemVariantRepository {
	return &itemVariantRepository{
		db: db,
	}
}

// struct method
func (r *itemVariantRepository) GetItemVariants(ctx context.Context, tenantID string, catalogItemID string, req dto.PaginationRequest) ([]model.ItemVariant, int64, error) {
	var iv []model.ItemVariant
	var total int64

	// query default
	query := r.db.WithContext(ctx).
		Model(model.ItemVariant{}).
		Preload("Tenant").
		Preload("Item").
		Where("tenant_id = ? AND item_id = ?", tenantID, catalogItemID) // Preload Item sesuaikan dengan model item_variants.go

	/*
		|--------------------------------------------------------------------------
		| Search
		|--------------------------------------------------------------------------
	*/
	if req.Search != "" {
		like := "%" + req.Search + "%"
		query = query.Where(`
			variant_name LIKE ?
			OR sku LIKE ?
			OR barcode LIKE ?
		`,
			like,
			like,
			like,
		)
	}

	/*
		|--------------------------------------------------------------------------
		| Count
		|--------------------------------------------------------------------------
	*/
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	/*
		|--------------------------------------------------------------------------
		| Sorting
		|--------------------------------------------------------------------------
		| Nanti akan kita whitelist.
		|--------------------------------------------------------------------------
	*/
	query = query.Order(req.Sort + " " + req.Order)

	/*
		|--------------------------------------------------------------------------
		| Pagination
		|--------------------------------------------------------------------------
	*/
	offset := (req.Page - 1) * req.Limit

	if err := query.
		Offset(offset).
		Limit(req.Limit).
		Find(&iv).Error; err != nil {

		return nil, 0, err
	}

	return iv, total, nil
}

func (r *itemVariantRepository) CountItemVariants(ctx context.Context, tenantID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(model.ItemVariant{}).
		Where("tenant_id = ?", tenantID).
		Count(&count).Error
	return count, err
}

func (r *itemVariantRepository) GetItemVariantByID(ctx context.Context, tenantID string, id string) (*model.ItemVariant, error) {
	var iv model.ItemVariant
	err := r.db.WithContext(ctx).
		Preload("Tenant").
		Preload("Item").
		Preload("Item.CatalogCategory"). // preload nested relation dari Item (untuk dapat kategori barang, misalnya snack, makanan, atau minuman)
		Where("tenant_id = ?", tenantID).First(&iv, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &iv, nil
}

func (r *itemVariantRepository) CreateItemVariant(ctx context.Context, iv *model.ItemVariant) error {
	return r.db.WithContext(ctx).Create(iv).Error
}

func (r *itemVariantRepository) UpdateItemVariant(ctx context.Context, tenantID string, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.ItemVariant{}).Where("id = ? AND tenant_id = ?", id, tenantID).Updates(updateMap).Error
}

func (r *itemVariantRepository) DeleteItemVariant(ctx context.Context, tenantID string, id string) error {
	return r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&model.ItemVariant{}).Error
}

// LOW STOCK SECTION
func (r *itemVariantRepository) GetLowStockItems(ctx context.Context, tenantID string) ([]dto.LowStockResponse, error) {
	var result []dto.LowStockResponse
	err := r.db.WithContext(ctx).Table("item_variants iv").
		Select(`iv.id AS item_variant_id,
		ci.name AS item_name,
		iv.sku,
		iv.variant_name,
		COALESCE(SUM(sm.qty),0) AS current_stock,
		iv.minimum_stock
	`).
		Joins(`JOIN catalog_items ci ON ci.id = iv.item_id`).
		Joins(`LEFT JOIN stock_movements sm ON sm.item_variant_id = iv.id`).
		Where(`iv.tenant_id = ? AND iv.is_active = TRUE`, tenantID).
		Group(`iv.id, ci.name, iv.sku, iv.variant_name, iv.minimum_stock`).
		Having(`COALESCE(SUM(sm.qty),0) <= iv.minimum_stock`).
		Order("ci.name ASC").
		Scan(&result).Error
	return result, err
}

func (r *itemVariantRepository) GetLowStockCount(ctx context.Context, tenantID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("item_variants iv").
		Select("iv.id").
		Joins(`JOIN catalog_items ci ON ci.id = iv.item_id`).
		Joins(`LEFT JOIN stock_movements sm ON sm.item_variant_id = iv.id`).
		Where(`iv.tenant_id = ? AND iv.is_active = TRUE`, tenantID).
		Group(`iv.id, iv.minimum_stock`).
		Having(`COALESCE(SUM(sm.qty),0) <= iv.minimum_stock`).
		Count(&count).Error
	return count, err
}
