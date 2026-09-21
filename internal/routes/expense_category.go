package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseCategoriesRoutes(rg *gin.RouterGroup, h *handler.ExpenseCategoryHandler) {
	// endpoint expense categories
	rg.GET("/expense-categories", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetExpenseCategory)
	// rg.GET("/expenses/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetExpenseByID)
	// rg.POST("/expenses", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin, constants.RoleCashier), h.CreateExpense)
	// rg.GET("/sales/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetSaleByID)
	// // untuk update status paid
	// rg.POST("/sales/:id/pay", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.PaySale)
}
