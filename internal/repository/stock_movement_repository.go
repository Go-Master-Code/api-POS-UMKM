package repository

import (
	"context"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// interface
type StockMovementRepository interface {
	CreateMovement(ctx context.Context, tx *gorm.DB, movement *model.StockMovement) error
	GetMovementByID(ctx context.Context, tenantID string, id string) (*model.StockMovement, error)
	GetMovementsByVariant(ctx context.Context, tenantID string, itemVariantID string) ([]model.StockMovement, error)
	GetCurrentStock(ctx context.Context, tenantID string, tx *gorm.DB, itemVariantID string) (float64, error)
	GetCurrentStockForUpdate(ctx context.Context, tenantID string, tx *gorm.DB, itemVariantID string) (float64, error)
}

// struct implementasi
type stockMovementRepository struct {
	db *gorm.DB
}

// constructor
func NewStockMovementRepository(db *gorm.DB) StockMovementRepository {
	return &stockMovementRepository{
		db: db,
	}
}

// struct method
func (r *stockMovementRepository) CreateMovement(ctx context.Context, tx *gorm.DB, movement *model.StockMovement) error {
	return tx.WithContext(ctx).Create(movement).Error // wajib di run pakai tx, bukan r.db karena merupakan alur 1 transaction
}

func (r *stockMovementRepository) GetMovementByID(ctx context.Context, tenantID string, id string) (*model.StockMovement, error) {
	var sm model.StockMovement
	err := r.db.WithContext(ctx).Preload("Tenant").Preload("ItemVariant").Preload("CreatedByUser").Where("tenant_id = ?", tenantID).First(&sm, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &sm, nil
}

func (r *stockMovementRepository) GetMovementsByVariant(ctx context.Context, tenantID string, itemVariantID string) ([]model.StockMovement, error) {
	var movements []model.StockMovement
	err := r.db.WithContext(ctx).
		Preload("Tenant").
		Preload("ItemVariant").
		Preload("CreatedByUser").
		Where("item_variant_id = ? AND tenant_id = ?", itemVariantID, tenantID).
		Order("created_at ASC").
		Find(&movements).Error

	if err != nil {
		return nil, err
	}

	return movements, nil
}

func (r *stockMovementRepository) GetCurrentStock(ctx context.Context, tenantID string, tx *gorm.DB, itemVariantID string) (float64, error) {
	var totalStock float64

	err := r.db.
		WithContext(ctx).
		Model(&model.StockMovement{}).
		Where("item_variant_id = ? AND tenant_id = ?", itemVariantID, tenantID).
		Select("COALESCE(SUM(qty), 0)").
		Scan(&totalStock).
		Error

	// coalesce -> jika belum ada movement sum(qty) maka akan return 0 bukan null

	if err != nil {
		return 0, err
	}

	return totalStock, nil
}

func (r *stockMovementRepository) GetCurrentStockForUpdate(ctx context.Context, tenantID string, tx *gorm.DB, itemVariantID string) (float64, error) {

	var totalStock float64

	// Lock item variant selama transaction berlangsung.
	var variant model.ItemVariant

	err := tx.
		WithContext(ctx).
		Clauses(clause.Locking{
			Strength: "UPDATE",
		}).
		Where(
			"id = ? AND tenant_id = ?",
			itemVariantID,
			tenantID,
		).
		First(&variant).
		Error

	if err != nil {
		return 0, err
	}

	// Setelah variant berhasil di-lock,
	// hitung stock dari seluruh movement.
	err = tx.
		WithContext(ctx).
		Model(&model.StockMovement{}).
		Where(
			"item_variant_id = ? AND tenant_id = ?",
			itemVariantID,
			tenantID,
		).
		Select("COALESCE(SUM(qty), 0)").
		Scan(&totalStock).
		Error

	if err != nil {
		return 0, err
	}

	return totalStock, nil
}
