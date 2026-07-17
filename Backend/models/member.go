package models

import "time"

type Member struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"unique;not null"`
	Discount  float64   `json:"discount" gorm:"default:10"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
