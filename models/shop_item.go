package models

import "time"

type ShopItem struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       int       `json:"price" gorm:"not null;default:0"`
	Icon        string    `json:"icon"`
	Category    string    `json:"category" gorm:"default:other"` // discount, cosmetic, service, other
	Stock       int       `json:"stock" gorm:"default:-1"`       // -1 = бесконечно
	IsAvailable bool      `json:"is_available" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
