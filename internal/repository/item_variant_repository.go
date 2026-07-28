package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type ItemVariantRepository interface {
	GetAllItemVariants(ctx context.Context, tenantID string) ([]model.ItemVariant, error)
	GetItemVariants(ctx context.Context, tenantID string, name string) ([]model.ItemVariant, error)
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
func (r *itemVariantRepository) GetAllItemVariants(ctx context.Context, tenantID string) ([]model.ItemVariant, error) {
	var variants []model.ItemVariant
	err := r.db.WithContext(ctx).Preload("Item").Where("tenant_id = ?", tenantID).Find(&variants).Error
	if err != nil {
		return nil, err
	}

	return variants, nil
}

func (r *itemVariantRepository) GetItemVariants(ctx context.Context, tenantID string, name string) ([]model.ItemVariant, error) {
	var iv []model.ItemVariant
	// query default
	query := r.db.WithContext(ctx).Preload("Tenant").Preload("Item").Where("tenant_id = ?", tenantID) // Preload Item sesuaikan dengan model item_variants.go

	if name != "" {
		// jika name nya tidak kosong, tambahkan query
		query = query.Where("variant_name LIKE ?", "%"+name+"%")
	}

	// find data by name
	err := query.Find(&iv).Error

	if err != nil {
		return nil, err
	}

	return iv, nil
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
