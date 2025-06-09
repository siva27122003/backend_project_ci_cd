package model

import (
	"time"

	"gorm.io/gorm"
)

type Farmer struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	FarmerID    int32          `gorm:"primaryKey;autoIncrement" json:"farmer_id"`
	UserID      int32          `json:"user_id"`                             // foreign key field (capital ID)
	User        User           `gorm:"foreignKey:UserID;references:Userid"` // association with explicit foreign key and reference
	DigitalId   string         `json:"digital_id"`
	LandHectare float32        `gorm:"not null" json:"land_in_hectares"`
}
