package models

import (
	"time"

	"gorm.io/gorm"
)

func (TransactionDetails) TableName() string {
	return "transaction_details"
}

type TransactionDetails struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Transaction Transactions `gorm:"foreignKey:TransactionID"`
}
