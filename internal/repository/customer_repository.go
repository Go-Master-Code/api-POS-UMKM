package repository

import (
	"context"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"

	"gorm.io/gorm"
)

// interface
type CustomerRepository interface {
	GetAllCustomesrPerTenant(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.Customer, int64, error)
	GetCustomerByID(ctx context.Context, tenantID string, id string) (*model.Customer, error)
	CreateCustomer(ctx context.Context, customer *model.Customer) error
	UpdateCustomer(ctx context.Context, tenantID string, id string, updateMap map[string]any) error
	DeleteCustomer(ctx context.Context, tenantID string, id string) error
}

// struct implementasi
type customerRepository struct {
	db *gorm.DB
}

// constructor
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) GetAllCustomesrPerTenant(ctx context.Context, tenantID string, req dto.PaginationRequest) ([]model.Customer, int64, error) {
	var customers []model.Customer
	var total int64 // untuk return value total data

	query := r.db.WithContext(ctx).Model(model.Customer{}).Preload("Tenant").Where("tenant_id = ?", tenantID)

	if req.Search != "" {
		query = query.Where("name LIKE ? OR phone LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	// hitung jumlah data sebelum pagination
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// sorting sementara
	query = query.Order(req.Sort + " " + req.Order)

	// Pagination
	offset := (req.Page - 1) * req.Limit

	// Limit + find data
	err = query.
		Offset(offset).
		Limit(req.Limit).
		Find(&customers).Error // sudah sekalian find data disini

	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (r *customerRepository) GetCustomerByID(ctx context.Context, tenantID string, id string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).
		Preload("Tenant").
		Where("tenant_id = ?", tenantID).
		First(&customer, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) CreateCustomer(ctx context.Context, customer *model.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

func (r *customerRepository) UpdateCustomer(ctx context.Context, tenantID string, id string, updateMap map[string]any) error {
	return r.db.WithContext(ctx).Model(model.Customer{}).Where("id = ? AND tenant_id = ?", id, tenantID).Updates(updateMap).Error
}

func (r *customerRepository) DeleteCustomer(ctx context.Context, tenantID string, id string) error {
	return r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&model.Customer{}).Error
}
