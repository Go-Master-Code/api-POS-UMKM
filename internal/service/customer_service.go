package service

import (
	"context"
	"math"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/repository"

	"github.com/google/uuid"
)

// interface
type CustomerService interface {
	GetAllCustomesrPerTenant(ctx context.Context, req dto.PaginationRequest) ([]dto.CustomerResponse, dto.PaginationResponse, error)
	GetCustomerByID(ctx context.Context, id string) (dto.CustomerResponse, error)
	CreateCustomer(ctx context.Context, req dto.CreateCustomerRequest) (dto.CustomerResponse, error)
	DeleteCustomer(ctx context.Context, id string) (dto.CustomerResponse, error)
	UpdateCustomer(ctx context.Context, id string, req dto.UpdateCustomerRequest) (dto.CustomerResponse, error)
}

// struct implementasi
type customerService struct {
	repo repository.CustomerRepository
}

// constructor
func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		repo: repo,
	}
}

// struct implementasi
func (s *customerService) GetAllCustomesrPerTenant(ctx context.Context, req dto.PaginationRequest) ([]dto.CustomerResponse, dto.PaginationResponse, error) {
	// get tenantID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	customers, total, err := s.repo.GetAllCustomesrPerTenant(ctx, tenantID, req)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	// convert model to dto
	customersDTO := helper.ConvertToDTOCustomerPlural(customers)

	// hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	// metadata pagination
	meta := dto.PaginationResponse{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return customersDTO, meta, nil
}

func (s *customerService) GetCustomerByID(ctx context.Context, id string) (dto.CustomerResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	customer, err := s.repo.GetCustomerByID(ctx, tenantID, id)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	// convert model to dto
	customerDTO := helper.ConvertToDTOCustomerSingle(customer)
	return customerDTO, nil
}

func (s *customerService) CreateCustomer(ctx context.Context, req dto.CreateCustomerRequest) (dto.CustomerResponse, error) {
	// ambil tenantID dari jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// parsing req ke model
	customer := model.Customer{
		ID:       uuid.NewString(),
		TenantID: tenantID,
		Name:     req.Name,
		Phone:    &req.Phone,
		IsActive: req.IsActive,
	}

	err := s.repo.CreateCustomer(ctx, &customer)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	// get customer by id untuk reload relasi
	newCustomer, err := s.repo.GetCustomerByID(ctx, tenantID, customer.ID)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	// convert model to dto
	customerDTO := helper.ConvertToDTOCustomerSingle(newCustomer)

	return customerDTO, nil
}

func (s *customerService) DeleteCustomer(ctx context.Context, id string) (dto.CustomerResponse, error) {
	// get tenantID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data by ID dulu untuk ditampilkan di response
	customer, err := s.repo.GetCustomerByID(ctx, tenantID, id)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	err = s.repo.DeleteCustomer(ctx, tenantID, id)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	// convert model to dto
	customerDTO := helper.ConvertToDTOCustomerSingle(customer)
	return customerDTO, nil
}

func (s *customerService) UpdateCustomer(ctx context.Context, id string, req dto.UpdateCustomerRequest) (dto.CustomerResponse, error) {
	// inisiasi update map
	var updateMap = map[string]any{}

	if req.Name != nil {
		updateMap["name"] = req.Name
	}
	if req.Phone != nil {
		updateMap["phone"] = req.Phone
	}
	if req.IsActive != nil {
		updateMap["is_active"] = req.IsActive
	}

	// get tenantID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// update ke repo
	err := s.repo.UpdateCustomer(ctx, tenantID, id, updateMap)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	// get data untuk preload relasi
	customer, err := s.repo.GetCustomerByID(ctx, tenantID, id)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	// convert model to dto
	customerDTO := helper.ConvertToDTOCustomerSingle(customer)

	return customerDTO, nil
}
