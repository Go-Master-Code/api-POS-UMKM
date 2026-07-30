package handler

import (
	"fmt"
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type UploadHandler struct {
	service service.UploadService
}

// constructor
func NewUploadHandler(service service.UploadService) *UploadHandler {
	return &UploadHandler{
		service: service,
	}
}

// struct method
func (h *UploadHandler) UploadLogo(c *gin.Context) {
	/*
		|--------------------------------------------------------------------------
		| Ambil file multipart
		|--------------------------------------------------------------------------
	*/
	file, err := c.FormFile("file")
	if err != nil {
		helper.ErrorResponse(c, "failed to get the file", err)
		return
	}

	/*
		|--------------------------------------------------------------------------
		| Upload melalui service
		|--------------------------------------------------------------------------
	*/
	fileName, err := h.service.UploadLogo(file)
	if err != nil {
		helper.ErrorResponse(c, "failed to upload the logo", err)
		return
	}

	// format url agar mudah dimasukkan ke pesan sukses
	url := fmt.Sprintf("/storage/logo/%s", fileName)

	// response berhasil
	helper.SuccessUploadLogo(c, "logo uploaded successfully", fileName, url)
}

func (h *UploadHandler) DeleteLogo(c *gin.Context) {
	// ambil nama file dari URL
	fileName := c.Param("fileName")
	if fileName == "" {
		helper.ErrorResponse(c, "file name is required", nil)
		return
	}

	err := h.service.DeleteLogo(fileName)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorDeleteData, err)
		return
	}

	helper.SuccessResponse(c, "logo deleted successfully", nil)
}
