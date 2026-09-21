package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// tidak ada interface, langsung struct implementasi
type ExpenseHandler struct {
	service service.ExpenseService
}

// constructor
func NewExpenseHandler(service service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

// struct method
func (h *ExpenseHandler) GetAllExpenses(c *gin.Context) {
	var query dto.GetAllExpenseQuery

	// parsing request body
	err := c.ShouldBindQuery(&query) // ambil parameter dari query string, bukan request body
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// pagination with helper
	helper.NormalizePagination(&query.Page, &query.Limit)

	expenses, total, err := h.service.GetAllExpenses(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	// jika sukses
	helper.SuccessGetAllExpensesPerTenant(c, expenses, int(total), query.Page, query.Limit)
}

func (h *ExpenseHandler) GetExpenseByID(c *gin.Context) {
	// get id from param
	id := c.Param("id")

	expense, err := h.service.GetExpenseByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	// jika sukses
	helper.SuccessResponse(c, constants.SuccessGetData, expense)
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var req dto.CreateExpenseRequest
	// parsing request body
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newExpense, err := h.service.CreateExpense(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	// success response
	helper.SuccessResponse(c, constants.SuccessCreateData, newExpense)
}
