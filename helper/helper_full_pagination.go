package helper

import (
	"strconv"
	"umkm-odod/internal/dto"

	"github.com/gin-gonic/gin"
)

// GetPagination membaca query parameter pagination dari request.
func GetPagination(c *gin.Context) dto.PaginationRequest {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	return dto.PaginationRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
		Sort:   c.DefaultQuery("sort", "created_at"),
		Order:  c.DefaultQuery("order", "desc"),
	}
}
