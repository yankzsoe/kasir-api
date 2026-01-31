package models

import (
	"time"

	"gorm.io/gorm"
)

func (Produk) TableName() string {
	return "products"
}

type Produk struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Price      int            `json:"price"`
	Stock      int            `json:"stock"`
	IsActive   bool           `gorm:"default:true" json:"is_active"`
	CategoryId *int           `json:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Category   Category       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
