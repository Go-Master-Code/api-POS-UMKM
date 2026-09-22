package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup, h *handler.ReportHandler) {
	rg.GET("/reports/sales", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSalesReport)
	rg.GET("/reports/purchase", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetPurchaseReport)
	rg.GET("/reports/stock", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetStockReport)
	// export report xlsx
	rg.GET("reports/sales/export", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportSalesReport)
	rg.GET("reports/purchase/export", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportPurchaseReport)
	rg.GET("reports/stock/export", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportStockHandler)
	// export sales invoice pdf (ukuran kertas printer 88 mm)
	rg.GET("/sales/:id/invoice", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin, constants.RoleCashier), h.ExportSalesInvoicePDF)
	// export report pdf
	rg.GET("reports/sales/pdf", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportSalesReportPDF)
	rg.GET("reports/expenses/pdf", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportExpensesReportPDF)
	rg.GET("reports/purchase/pdf", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportPurchaseReportPDF)
	rg.GET("reports/stock/pdf", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportStockReportPDF)
	rg.GET("reports/stock-card/pdf/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.ExportStockCardPDF)
}
