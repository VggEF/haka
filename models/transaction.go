package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"not null;index"`
	Amount    int       `json:"amount"`  // положительный - начисление, отрицательный - списание
	Type      string    `json:"type"`    // "earn", "spend", "bonus", "refund"
	Reason    string    `json:"reason"`  // "daily_bonus", "achievement", "purchase", etc.
	ItemID    *int      `json:"item_id"` // ID товара если это покупка
	CreatedAt time.Time `json:"created_at"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}
