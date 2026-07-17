package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Username  string    `json:"username" gorm:"unique; not null; size:255"`
	Password  string    `json:"password" gorm:"not null; size:255"`
	Role      string    `json:"role" gorm:"not null; size:50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
