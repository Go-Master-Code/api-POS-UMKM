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
type RoleService interface {
	GetRolesByTenant(ctx context.Context, req dto.PaginationRequest) ([]dto.RoleResponse, dto.PaginationResponse, error)
	GetRoleByID(ctx context.Context, id string) (dto.RoleResponse, error)
	CreateRole(ctx context.Context, req dto.CreateRoleRequest) (dto.RoleResponse, error)
	UpdateRole(ctx context.Context, id string, req dto.UpdateRoleRequest) (dto.RoleResponse, error)
	DeleteRole(ctx context.Context, id string) (dto.RoleResponse, error)
}

// struct impelemtasi
type roleService struct {
	repo repository.RoleRepository
}

// constructor
func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{
		repo: repo,
	}
}

// struct method
func (s *roleService) GetRolesByTenant(ctx context.Context, req dto.PaginationRequest) ([]dto.RoleResponse, dto.PaginationResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	roles, total, err := s.repo.GetRolesByTenant(ctx, tenantID, req)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	// convert model to dto
	rolesDTO := helper.ConvertToDTORolePlural(roles)

	// hitung total halaman
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	// metadata pagination
	meta := dto.PaginationResponse{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return rolesDTO, meta, nil
}

func (s *roleService) GetRoleByID(ctx context.Context, id string) (dto.RoleResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	role, err := s.repo.GetRoleByID(ctx, tenantID, id)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(role)
	return roleDTO, nil
}

func (s *roleService) CreateRole(ctx context.Context, req dto.CreateRoleRequest) (dto.RoleResponse, error) {
	// ambil tenantID dari context -> cek file middleware/auth_required.go
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// payload sementara
	// tenantID := "f27e441f-5385-4b8d-b2e2-88b8615a4634"

	// convert dto to model
	role := model.Role{
		TenantID: tenantID,
		ID:       uuid.NewString(),
		Name:     req.Name,
	}

	err := s.repo.CreateRole(ctx, &role)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// get role by id (preload tenant) agar nama tenant muncul
	newRole, err := s.repo.GetRoleByID(ctx, tenantID, role.ID)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(newRole)
	return roleDTO, nil
}

func (s *roleService) UpdateRole(ctx context.Context, id string, req dto.UpdateRoleRequest) (dto.RoleResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// mapping update data ke map
	var updateMap = map[string]any{}

	if req.Name != nil {
		updateMap["name"] = req.Name
	}

	// update data
	err := s.repo.UpdateRole(ctx, tenantID, id, updateMap)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// get data yang sudah diupdate untuk return value
	updateRole, err := s.repo.GetRoleByID(ctx, tenantID, id)

	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(updateRole)
	return roleDTO, nil
}

func (s *roleService) DeleteRole(ctx context.Context, id string) (dto.RoleResponse, error) {
	// get tenant ID from jwt
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// get data dulu sebelum di delete
	role, err := s.repo.GetRoleByID(ctx, tenantID, id)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// delete data
	err = s.repo.DeleteRole(ctx, tenantID, id)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	// convert model to dto
	roleDTO := helper.ConvertToDTORoleSingle(role)
	return roleDTO, nil
}
