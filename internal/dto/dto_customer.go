package dto

import "time"

// response
type CustomerResponse struct {
	ID         string    `json:"id"`
	TenantID   string    `json:"tenant_id"`
	TenantName string    `json:"tenant_name"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}

// create request
type CreateCustomerRequest struct {
	// ID          string  `json:"id" binding:"required,uuid"` tidak usah karena uuid di generate backend, bukan dari request body
	// TenantID    string  `json:"tenant_id" binding:"required,uuid"` jangan dari create request, berasal dari jwt harusnya
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Phone    string `json:"phone" binding:"omitempty,min=8,max=30"`
	IsActive bool   `json:"is_active"`
}

// update request
type UpdateCustomerRequest struct {
	Name     *string `json:"name" binding:"omitempty,min=3,max=100"`
	Phone    *string `json:"phone" binding:"omitempty,min=8,max=30"`
	IsActive *bool   `json:"is_active"` // update request var bool tidak perlu omitempty
}
