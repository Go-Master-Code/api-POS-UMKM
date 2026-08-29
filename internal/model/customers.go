package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID       string  `json:"id" gorm:"type:char(36);primaryKey"`
	TenantID string  `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant   Tenant  `json:"-" gorm:"foreignKey:TenantID"`
	Name     string  `json:"name" gorm:"type:varchar(100);not null"`
	Phone    *string `json:"phone" gorm:"type:varchar(30)"`
	IsActive bool    `json:"is_active" gorm:"default:true"`
	// Relasi ke Sales
	Sale      []Sale         `json:"-" gorm:"foreignKey:CustomerID"` // Relasi 1 to many - 1 Customer bisa membuat beberapa sales
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Customer) TableName() string {
	return "customers"
}
