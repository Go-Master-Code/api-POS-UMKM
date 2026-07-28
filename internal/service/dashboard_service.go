package service

import (
	"context"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/repository"
)

// interface
type DashboardService interface {
	GetSummary(ctx context.Context) (dto.DashBoardSummaryResponse, error)
	GetDailySalesChart(ctx context.Context) ([]dto.DailySalesChartResponse, error)
	GetDailyPurchaseChart(ctx context.Context) ([]dto.DailyPurchaseChartResponse, error)
	GetTopSellingProducts(ctx context.Context) ([]dto.TopSellingProductsResponse, error)
	GetRecentSales(ctx context.Context) ([]dto.RecentSalesResponse, error)
}

// struct implementasi
type dashboardService struct {
	dashboardRepo     repository.DashboardRepository
	stockMovementRepo repository.StockMovementRepository
	itemVariantRepo   repository.ItemVariantRepository
	catalogItemRepo   repository.CatalogItemRepository
	supplierRepo      repository.SupplierRepository
}

// constructor
func NewDashboardService(dashboardRepo repository.DashboardRepository, stockMovementRepo repository.StockMovementRepository, itemVariantRepo repository.ItemVariantRepository, catalogItemRepo repository.CatalogItemRepository, supplierRepo repository.SupplierRepository) DashboardService {
	return &dashboardService{
		dashboardRepo:     dashboardRepo,
		stockMovementRepo: stockMovementRepo,
		itemVariantRepo:   itemVariantRepo,
		catalogItemRepo:   catalogItemRepo,
		supplierRepo:      supplierRepo,
	}
}

// struct method
func (s *dashboardService) GetSummary(ctx context.Context) (dto.DashBoardSummaryResponse, error) {
	// get tenantID by context
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	// =========================
	// ==========SALES==========
	// =========================
	totalSales, totalTransactions, err := s.dashboardRepo.GetTodaySales(ctx, tenantID)
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// =============================
	// ==========LOW STOCK==========
	// =============================
	lowStockCount, err := s.itemVariantRepo.GetLowStockCount(ctx, tenantID)
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// ============================
	// ==========PURCHASE==========
	// ============================
	totalPurchases, totalPurchaseTransactions, err := s.dashboardRepo.GetTodayPurchases(ctx, tenantID)
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// =============================
	// ==========SUPPLIERS==========
	// =============================
	suppliers, err := s.supplierRepo.GetSuppliers(ctx, tenantID, "")
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// convert int ke int64
	totalSuppliers := int64(len(suppliers))

	// ==============================================
	// ==========CATALOG ITEM (JUMLAH ITEM)==========
	// ==============================================
	catalogItems, err := s.catalogItemRepo.GetCatalogItems(ctx, tenantID, "")
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// convert int ke int64
	totalItems := int64(len(catalogItems))

	// =================================
	// ==========ITEM VARIANTS==========
	// =================================
	totalVariants, err := s.itemVariantRepo.CountItemVariants(ctx, tenantID)
	if err != nil {
		return dto.DashBoardSummaryResponse{}, err
	}

	// =================================
	// MAPPING SEMUA HASIL QUERY KE DTO
	// =================================
	dashboardSummary := dto.DashBoardSummaryResponse{
		TodaySales:        totalSales,
		TodayTransactions: totalTransactions,
		LowStockCount:     lowStockCount,
		TodayPurchase:     totalPurchases,
		// TodayProfit:              totalSales - totalPurchases, dihitung nanti
		TodayPurchaseTransaction: totalPurchaseTransactions,
		TotalSuppliers:           totalSuppliers,
		TotalItems:               totalItems,
		TotalVariants:            totalVariants,
	}

	return dashboardSummary, nil
}

func (s *dashboardService) GetDailySalesChart(ctx context.Context) ([]dto.DailySalesChartResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	dailySalesChart, err := s.dashboardRepo.GetDailySalesChart(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return dailySalesChart, nil
}

func (s *dashboardService) GetDailyPurchaseChart(ctx context.Context) ([]dto.DailyPurchaseChartResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	dailyPurchaseChart, err := s.dashboardRepo.GetDailyPurchaseChart(ctx, tenantID)

	if err != nil {
		return nil, err
	}

	return dailyPurchaseChart, nil
}

func (s *dashboardService) GetTopSellingProducts(ctx context.Context) ([]dto.TopSellingProductsResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	topSellingProducts, err := s.dashboardRepo.GetTopSellingProducts(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return topSellingProducts, nil
}

func (s *dashboardService) GetRecentSales(ctx context.Context) ([]dto.RecentSalesResponse, error) {
	// get tenantID from ctx
	tenantID := ctx.Value(constants.ContextTenantID).(string)

	const limit = 5 // dashboard hanya membutuhkan 5 trx terakhir

	recentSales, err := s.dashboardRepo.GetRecentSales(ctx, tenantID, limit)
	if err != nil {
		return nil, err
	}

	return recentSales, nil
}
