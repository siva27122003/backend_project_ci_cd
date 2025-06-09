package model

import (
	"time"

	"gorm.io/gorm"
)

type Bidding struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Bidid       int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	CommodityID int32          `json:"commodity_id"`
	Commodity   Commodity      `gorm:"foreignkey:CommodityID;references:CommodityID"`
	Userid      int32          `json:"user_id"`
	User        User           `gorm:"foreignkey:Userid;references:Userid"`
	BidAmount   float32        `gorm:"not null" json:"bid_amount"`
	BidStatus   string         `gorm:"default:pending" json:"Status"`
}
