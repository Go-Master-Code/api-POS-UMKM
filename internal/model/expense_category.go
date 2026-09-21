package model

import (
	"time"

	"gorm.io/gorm"
)

type ExpenseCategory struct {
	ID          string         `json:"id" gorm:"type:char(36);primaryKey"` // format jika uuid
	TenantID    string         `json:"tenant_id" gorm:"type:char(36);not null;index"`
	Tenant      Tenant         `json:"-" gorm:"foreignKey:TenantID"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null"`
	Description string         `json:"description" gorm:"type:varchar(255);"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (ExpenseCategory) TableName() string {
	return "expense_categories"
}
