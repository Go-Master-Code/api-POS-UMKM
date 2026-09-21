package service

import (
	"context"
	"math"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"
)

// interface
type ExpenseCategoryService interface {
	GetExpenseCategory(ctx context.Context, req dto.PaginationRequest) ([]dto.ExpenseCategoryResponse, dto.PaginationResponse, error)
}

// struct implementasi
type expenseCategoryService struct {
	repo repository.ExpenseCategoryRepository
}

// constructur
func NewExpenseCategoryService(repo repository.ExpenseCategoryRepository) ExpenseCategoryService {
	return &expenseCategoryService{
		repo: repo,
	}
}

// struct method
func (s *expenseCategoryService) GetExpenseCategory(ctx context.Context, req dto.PaginationRequest) ([]dto.ExpenseCategoryResponse, dto.PaginationResponse, error) {
	// get tenantID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	expenseCategories, total, err := s.repo.GetExpenseCategory(ctx, tenantID, req)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	// convert model to dto
	ecDTO := helper.ConvertTODTOExpenseCategoryPlural(expenseCategories)

	// hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	// metadata pagination
	meta := dto.PaginationResponse{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return ecDTO, meta, nil
}
