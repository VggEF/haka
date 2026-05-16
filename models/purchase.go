package models

import "time"

type Purchase struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	UserID        int       `json:"user_id" gorm:"not null;index"`
	ItemID        int       `json:"item_id" gorm:"not null"`
	TransactionID int       `json:"transaction_id"`
	Price         int       `json:"price"`
	CreatedAt     time.Time `json:"created_at"`

	User User     `json:"user" gorm:"foreignKey:UserID"`
	Item ShopItem `json:"item" gorm:"foreignKey:ItemID"`
}
