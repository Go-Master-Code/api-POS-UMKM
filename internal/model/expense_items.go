package model

import "time"

type ExpenseItem struct {
	ID                string          `json:"id" gorm:"type:char(36);primaryKey"`
	ExpenseID         string          `json:"expense_id" gorm:"type:char(36);not null;index"`
	Expenses          Expenses        `json:"-" gorm:"foreignKey:ExpenseID"` // reverse relation to model sale
	ExpenseCategoryID string          `json:"expense_category_id" gorm:"type:char(36);not null;index"`
	ExpenseCategory   ExpenseCategory `json:"-" gorm:"foreignKey:ExpenseCategoryID"`
	Description       string          `json:"description" gorm:"type:varchar(255)"`
	Qty               float64         `json:"qty" gorm:"type:decimal(18,2);not null"`
	Unit              string          `json:"unit" gorm:"type:varchar(30);not null"`
	UnitPrice         float64         `json:"unit_price" gorm:"type:decimal(18,2);not null"`
	Subtotal          float64         `json:"subtotal" gorm:"type:decimal(18,2);not null"`
	CreatedAt         time.Time       `gorm:"column:created_at;autoCreateTime"`
}

func (ExpenseItem) TableName() string {
	return "expense_items"
}
