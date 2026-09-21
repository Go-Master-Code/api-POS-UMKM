package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type ExpenseCategoryHandler struct {
	service service.ExpenseCategoryService
}

// constructor
func NewExpenseCategoryHandler(service service.ExpenseCategoryService) *ExpenseCategoryHandler {
	return &ExpenseCategoryHandler{
		service: service,
	}
}

// struct method
func (h *ExpenseCategoryHandler) GetExpenseCategory(c *gin.Context) {
	// ambil pagination dari query, sudah berisi juga keyword untuk search data dari table
	req := helper.GetPagination(c)

	ec, meta, err := h.service.GetExpenseCategory(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessPaginationResponse(c, constants.SuccessGetData, ec, meta)
}
