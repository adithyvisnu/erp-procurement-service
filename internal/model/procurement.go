package model

import "time"

type Supplier struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:200;not null"`
	Email     string `gorm:"size:200"`
	Phone     string `gorm:"size:50"`
	CreatedAt time.Time
}

type Item struct {
	ID        uint    `gorm:"primaryKey"`
	SKU       string  `gorm:"size:64;uniqueIndex;not null"`
	Name      string  `gorm:"size:200;not null"`
	UnitPrice float64 `gorm:"type:numeric(14,2);not null;default:0"`
	CreatedAt time.Time
}

type PurchaseRequest struct {
	ID        uint   `gorm:"primaryKey"`
	Requester string `gorm:"size:100;not null"`
	Status    string `gorm:"size:32;index;not null;default:PENDING"`
	Items     []PurchaseRequestItem
	CreatedAt time.Time
}

type PurchaseRequestItem struct {
	PurchaseRequestID uint `gorm:"primaryKey;index"`
	ItemID            uint `gorm:"primaryKey;index"`
	Qty               int  `gorm:"not null"`
}
