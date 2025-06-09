package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Userid    int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"not null" json:"user_name"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Phone     string         `gorm:"not null" json:"phone_number"`
	Password  string         `gorm:"not null" json:"password"`
	Role      string         `gorm:"not null" json:"role"`
	Location  string         `gorm:"not null" json:"location"`
}
