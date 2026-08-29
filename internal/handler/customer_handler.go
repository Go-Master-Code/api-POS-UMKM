package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type CustomerHandler struct {
	service service.CustomerService
}

// constructor
func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		service: service,
	}
}

// struct method
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	// ambil pagination dari query, sudah berisi juga keyword untuk search data dari table
	req := helper.GetPagination(c)

	customerDTO, meta, err := h.service.GetAllCustomesrPerTenant(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessPaginationResponse(c, constants.SuccessGetData, customerDTO, meta)
}

func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
	// ambil param id
	id := c.Param("id")

	customer, err := h.service.GetCustomerByID(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessGetData, customer)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req dto.CreateCustomerRequest
	// parsing request body
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	customer, err := h.service.CreateCustomer(c.Request.Context(), req)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorCreateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessCreateData, customer)
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	var req dto.UpdateCustomerRequest

	// parsing request body
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")
	newCustomer, err := h.service.UpdateCustomer(c.Request.Context(), id, req)

	if err != nil {
		helper.ErrorResponse(c, constants.ErrorUpdateData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessUpdateData, newCustomer)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	// ambil param id
	id := c.Param("id")

	customer, err := h.service.DeleteCustomer(c.Request.Context(), id)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorDeleteData, err)
		return
	}

	helper.SuccessResponse(c, constants.SuccessDeleteData, customer)
}
