package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type RoleRepository interface {
	GetRolesByTenant(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.Role, int64, error)
	GetRoleByID(ctx context.Context, tenantID string, id string) (*model.Role, error)
	CreateRole(ctx context.Context, role *model.Role) error
	UpdateRole(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	DeleteRole(ctx context.Context, tenantID string, id string) error
}

// struct implementasi
type roleRepository struct {
	db *gorm.DB
}

// constructor
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

// struct method
func (r *roleRepository) GetRolesByTenant(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64 // untuk return value total data

	// query search sementara tanpa name
	query := r.db.WithContext(ctx).Model(&model.Role{}).Preload("Tenant").Where("tenant_id = ?", tenantID)

	if req.Search != "" {
		query = query.Where("name LIKE ?", "%"+req.Search+"%")
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

	// Limit + find data
	err = query.
		Offset(offset).
		Limit(req.Limit).
		Find(&roles).Error // sudah sekalian find data disini

	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (r *roleRepository) GetRoleByID(ctx context.Context, tenantID string, id string) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).Preload("Tenant").Where("tenant_id = ?", tenantID).First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) CreateRole(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) UpdateRole(ctx context.Context, tenantID string, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.Role{}).Where("id = ? AND tenant_id = ?", id, tenantID).Updates(updateMap).Error
}

func (r *roleRepository) DeleteRole(ctx context.Context, tenantID string, id string) error {
	return r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&model.Role{}).Error
}
