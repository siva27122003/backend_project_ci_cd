package model

import (
	"time"

	"gorm.io/gorm"
)

type Commodity struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	CommodityID  int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductName  string         `gorm:"not null" json:"product_name"`
	FarmerID     int32          `json:"farmer_id"`
	Farmer       Farmer         `gorm:"foreignkey : FarmerID;references:FarmerID"`
	Quantity     int32          `gorm:"not null" json:"quantity"`
	BasePrice    float64        `gorm:"not null" json:"base_price"`
	Availability bool           `gorm:"default:true" json:"availability"`
	CategoryID   int32          `json:"category_id"`
	Category     Category       `gorm:"foreignkey : CategoryID;references:ID"`
}
