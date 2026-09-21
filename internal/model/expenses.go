package model

import (
	"time"
)

type Expenses struct {
	ID            string        `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID      string        `json:"tenant_id" gorm:"type:char(36);not null;index"`
	TotalAmount   float64       `json:"total_amount" gorm:"type:decimal(18,2);not null"`
	PaymentMethod string        `json:"payment_method" gorm:"type:varchar(20);not null"`
	ExpenseNumber string        `json:"expense_number" gorm:"type:varchar(30);not null"`
	Notes         string        `json:"notes" gorm:"type:text"`
	CreatedBy     string        `json:"created_by" gorm:"type:char(36);not null;index"`
	CreatedAt     time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time     `gorm:"column:updated_at;autoUpdateTime"`
	ExpenseItems  []ExpenseItem `json:"-" gorm:"foreignKey:ExpenseID"` // one expense has many expense items, relasi detail transaksi. ExpenseID diambil dari model expense_items
	// RELATIONS
	Tenant Tenant `json:"-" gorm:"foreignKey:TenantID"`
	User   User   `json:"-" gorm:"foreignKey:CreatedBy"`
}
