package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(rg *gin.RouterGroup, h *handler.CustomerHandler) {
	// endpoint customer
	rg.GET("/customers", h.GetCustomers)
	rg.GET("/customers/:id", h.GetCustomerByID)
	rg.POST("/customers", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateCustomer)
	rg.PUT("/customers/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateCustomer)
	rg.DELETE("/customers/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.DeleteCustomer)
}
