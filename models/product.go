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
	Nama       string         `json:"nama"`
	Harga      int            `json:"harga"`
	Stok       int            `json:"stok"`
	IsActive   bool           `gorm:"default:true" json:"is_active"`
	CategoryId *int           `json:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Category   Category       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
