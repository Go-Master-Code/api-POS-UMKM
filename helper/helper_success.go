package helper

import (
	"net/http"
	"umkm-odod/internal/dto"

	"github.com/gin-gonic/gin"
)

// definisikan struct global untuk success message
type AllSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type UploadLogoResponse struct {
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type AllSuccessPagination struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    any                    `json:"data"`
	Meta    dto.PaginationResponse `json:"meta"`
}

type GetAllSalesOrPurchaseOrExpensePerTenantSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Total   int    `json:"total"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

type LoginSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Token   string `json:"token"`
}

func SuccessResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, AllSuccess{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// success response yang sudah dilengkapi dengan page, limit, search, sort, order
func SuccessPaginationResponse(c *gin.Context, message string, data any, meta dto.PaginationResponse) {
	c.JSON(http.StatusOK, AllSuccessPagination{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func SuccessLogin(c *gin.Context, data any, token string) {
	c.JSON(http.StatusOK, LoginSuccess{
		Code:    http.StatusOK,
		Message: "login success",
		Data:    data,
		Token:   token,
	})
}

func SuccessGetAllSalesPerTenant(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchaseOrExpensePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success get all sales data",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

func SuccessGetAllExpensesPerTenant(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchaseOrExpensePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success get all expense data",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

func SuccessGetAllPurchasesPerTenant(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchaseOrExpensePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success get all purchases data",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

func SuccessGetAllPurchaseReturnPerTenant(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchaseOrExpensePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success get all purchase return data",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

func SuccessGenerateStockReport(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchaseOrExpensePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success generate stock report",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

func SuccessGetAllLogsPerTenant(c *gin.Context, data any, total int, page int, limit int) {
	c.JSON(http.StatusOK, GetAllSalesOrPurchaseOrExpensePerTenantSuccess{
		Code:    http.StatusOK,
		Message: "success get all activity logs",
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	})
}

// success upload file
func SuccessUploadLogo(c *gin.Context, message, fileName, url string) {
	c.JSON(http.StatusOK, AllSuccess{
		Code:    http.StatusOK,
		Message: message,
		Data: UploadLogoResponse{
			FileName: fileName,
			URL:      url,
		},
	})
}
