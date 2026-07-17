package models

import "time"

type Product struct {
	ID         uint       `json:"id" gorm:"primarykey"`
	Name       string     `json:"name" gorm:"not null;index"`
	Category   string     `json:"category" gorm:"not null;index"`
	Price      float64    `json:"price" gorm:"not null"`
	Stock      int        `json:"stock" gorm:"not null"`
	ExpiryDate *time.Time `json:"expiry_date"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
