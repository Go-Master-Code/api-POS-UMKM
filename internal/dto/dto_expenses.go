package dto

import (
	"time"
)

// struct response
type ExpensesResponse struct {
	ID            string                `json:"id"`
	TenantID      string                `json:"tenant_id"`
	TenantName    string                `json:"tenant_name"`
	TotalAmount   float64               `json:"total_amount"`
	PaymentMethod string                `json:"payment_method"`
	ExpenseNumber string                `json:"expense_number"`
	Notes         string                `json:"notes"`
	CreatedBy     string                `json:"created_by"`
	CreatedByUser string                `json:"created_by_user"`
	CreatedAt     time.Time             `json:"created_at"`
	Items         []ExpenseItemResponse `json:"items"` // tampilkan expense item sebagai nested slice
}

// create request
type CreateExpenseRequest struct {
	// TotalAmount   float64 `json:"total_amount" binding:"required,gte=0"`
	PaymentMethod string                           `json:"payment_method" binding:"required,oneof=CASH QRIS TRANSFER DEBIT KREDIT"`
	Notes         string                           `json:"notes" binding:"omitempty,max=500"`
	Items         []CreateExpenseItemDetailRequest `json:"items" binding:"required,min=1,dive"` // tabel detil sales berisi item yang dijual
}

// detil expense item pada transaksi expense
type CreateExpenseItemDetailRequest struct {
	ExpenseCategoryID string  `json:"expense_category_id" binding:"required,uuid"`
	Description       string  `json:"description" binding:"required"`
	Qty               float64 `json:"qty" binding:"required,gte=0"`
	Unit              string  `json:"unit" binding:"required"`
	UnitPrice         float64 `json:"unit_price" binding:"required,gte=0"`
}

// response expense item
type ExpenseItemResponse struct {
	ID                  string    `json:"id"`
	ExpenseID           string    `json:"expense_id"`
	ExpenseCategoryID   string    `json:"expense_category_id"`
	ExpenseCategoryName string    `json:"expense_category_name"`
	Description         string    `json:"description"`
	Qty                 float64   `json:"qty"`
	Unit                string    `json:"unit"`
	UnitPrice           float64   `json:"unit_price"`
	Subtotal            float64   `json:"subtotal"`
	CreatedAt           time.Time `json:"created_at"`
}

// query params get all sales
type GetAllExpenseQuery struct {
	Page          int    `form:"page"`           // nomor halaman
	Limit         int    `form:"limit"`          // jumlah data per halaman
	Search        string `form:"search"`         // untuk search total_amount
	PaymentMethod string `form:"payment_method"` // filter payment method
	Sort          string `form:"sort"`
	Order         string `form:"order"`
}
