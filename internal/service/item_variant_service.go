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
type ItemVariantService interface {
	GetItemVariants(ctx context.Context, catalogItemID string, req dto.PaginationRequest) ([]dto.ItemVariantResponse, dto.PaginationResponse, error)
	GetItemVariantByID(ctx context.Context, id string) (dto.ItemVariantResponse, error)
	GetLowStockItem(ctx context.Context) ([]dto.LowStockResponse, error)
	CountLowStockItem(ctx context.Context) (int64, error)
	CreateItemVariant(ctx context.Context, req dto.CreateItemVariantRequest) (dto.ItemVariantResponse, error)
	UpdateItemVariant(ctx context.Context, id string, req dto.UpdateItemVariantRequest) (dto.ItemVariantResponse, error)
	DeleteItemVariant(ctx context.Context, id string) (dto.ItemVariantResponse, error)
}

// struct implementasi
type itemVariantService struct {
	itemVariantRepo   repository.ItemVariantRepository
	stockMovementRepo repository.StockMovementRepository
}

// constructor
func NewItemVariantService(itemVariantRepo repository.ItemVariantRepository, stockMovementRepo repository.StockMovementRepository) ItemVariantService {
	return &itemVariantService{
		itemVariantRepo:   itemVariantRepo,
		stockMovementRepo: stockMovementRepo,
	}
}

// struct method
func (s *itemVariantService) GetItemVariants(ctx context.Context, catalogItemID string, req dto.PaginationRequest) ([]dto.ItemVariantResponse, dto.PaginationResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	iv, total, err := s.itemVariantRepo.GetItemVariants(ctx, tenantID, catalogItemID, req)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantPlural(iv)

	// hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	// metadata pagination
	meta := dto.PaginationResponse{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return ivDTO, meta, nil
}

func (s *itemVariantService) GetItemVariantByID(ctx context.Context, id string) (dto.ItemVariantResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	iv, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(iv)
	return ivDTO, err
}

func (s *itemVariantService) GetLowStockItem(ctx context.Context) ([]dto.LowStockResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	lowStockItems, err := s.itemVariantRepo.GetLowStockItems(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return lowStockItems, nil
}

func (s *itemVariantService) CountLowStockItem(ctx context.Context) (int64, error) {
	tenantID := ctx.Value(constants.ContextTenantID).(string)
	count, err := s.itemVariantRepo.GetLowStockCount(ctx, tenantID)
	return count, err
}

func (s *itemVariantService) CreateItemVariant(ctx context.Context, req dto.CreateItemVariantRequest) (dto.ItemVariantResponse, error) {
	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// parsing dto ke model
	iv := model.ItemVariant{
		ID:          uuid.NewString(),
		TenantID:    tenantID,
		ItemID:      req.ItemID,
		SKU:         req.SKU,
		Barcode:     req.Barcode,
		CostPrice:   req.CostPrice,
		VariantName: req.VariantName,
		IsActive:    req.IsActive,
	}

	err := s.itemVariantRepo.CreateItemVariant(ctx, &iv)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// get data by id
	newIV, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, iv.ID)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(newIV)
	return ivDTO, nil
}

func (s *itemVariantService) UpdateItemVariant(ctx context.Context, id string, req dto.UpdateItemVariantRequest) (dto.ItemVariantResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// mapping req ke map
	var updateMap = map[string]any{}

	if req.ItemID != nil {
		updateMap["item_id"] = req.ItemID
	}
	if req.SKU != nil {
		updateMap["sku"] = req.SKU
	}
	if req.Barcode != nil {
		updateMap["barcode"] = req.Barcode
	}
	if req.VariantName != nil {
		updateMap["variant_name"] = req.VariantName
	}
	if req.CostPrice != nil {
		updateMap["cost_price"] = req.CostPrice
	}
	if req.IsActive != nil {
		updateMap["is_active"] = req.IsActive
	}

	err := s.itemVariantRepo.UpdateItemVariant(ctx, tenantID, id, updateMap)

	// get data by id
	iv, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(iv)

	return ivDTO, nil
}

func (s *itemVariantService) DeleteItemVariant(ctx context.Context, id string) (dto.ItemVariantResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get id dulu untuk response api
	iv, err := s.itemVariantRepo.GetItemVariantByID(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// delete data
	err = s.itemVariantRepo.DeleteItemVariant(ctx, tenantID, id)
	if err != nil {
		return dto.ItemVariantResponse{}, err
	}

	// convert model to dto
	ivDTO := helper.ConvertToDTOItemVariantSingle(iv)
	return ivDTO, nil
}
