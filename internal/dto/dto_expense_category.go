package dto

import (
	"time"
)

// dto response
type ExpenseCategoryResponse struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	TenantName  string    `json:"tenant_name"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

// create request
type CreateExpenseCategoryRequest struct {
	Name        string `json:"name" binding:"required,min=5,max=100"`
	Description string `json:"description" binding:"omitempty,max=255"`
	IsActive    bool   `json:"is_active"`
}

// update request
