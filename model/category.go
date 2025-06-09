package model

type Category struct {
	ID           int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryName string `gorm:"not null;unique" json:"category_name"`
}
