package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterItemVariantRoutes(rg *gin.RouterGroup, h *handler.ItemVariantHandler) {
	// ============================
	// Variant milik Catalog Item
	// ============================
	rg.GET("/item_variants", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetItemVariants)
	rg.POST("/item_variants", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.CreateItemVariant)

	// ============================
	// Variant individual
	// ============================
	rg.GET("/item_variants/:id", h.GetItemVariantByID)
	rg.GET("/item_variants/low-stock", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.GetLowStockItem)
	rg.PUT("/item_variants/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.UpdateItemVariant)
	rg.DELETE("/item_variants/:id", middleware.AuthRole(constants.RoleOwner, constants.RoleAdmin), h.DeleteItemVariant)
}
