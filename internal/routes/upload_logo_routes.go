package routes

import (
	"umkm-odod/internal/constants"
	"umkm-odod/internal/handler"
	"umkm-odod/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUploadLogoRoutes(rg *gin.RouterGroup, h *handler.UploadHandler) {
	// endpoint upload logo
	rg.POST("/upload/logo", middleware.AuthRole(constants.RoleAdmin, constants.RoleOwner), h.UploadLogo)
	rg.DELETE("/upload/logo/:fileName", middleware.AuthRole(constants.RoleAdmin, constants.RoleOwner), h.DeleteLogo)
}
