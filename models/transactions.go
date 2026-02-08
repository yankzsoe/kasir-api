package models

import (
	"time"

	"gorm.io/gorm"
)

func (Transactions) TableName() string {
	return "transactions"
}

type Transactions struct {
	ID          int `gorm:"primaryKey"`
	TotalAmount int `json:"total_amount"`

	Details []TransactionDetails `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE;" json:"details"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
