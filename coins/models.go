package coins

import "time"

// ShopItem модель товара в магазине
type ShopItem struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Price       int       `json:"price" gorm:"not null;default:0"`
	Icon        string    `json:"icon" gorm:"type:varchar(50)"`
	Category    string    `json:"category" gorm:"type:varchar(50);default:'other'"`
	Stock       int       `json:"stock" gorm:"default:-1"`
	IsAvailable bool      `json:"is_available" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (ShopItem) TableName() string {
	return "shop_items"
}

// Transaction модель транзакции коинов
type Transaction struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"not null;index"`
	Amount    int       `json:"amount"`
	Type      string    `json:"type" gorm:"type:varchar(50)"` // earn, spend, bonus, refund
	Reason    string    `json:"reason" gorm:"type:varchar(255)"`
	ItemID    *int      `json:"item_id" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (Transaction) TableName() string {
	return "transactions"
}

// Purchase модель покупки
type Purchase struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        int       `json:"user_id" gorm:"not null;index"`
	ItemID        int       `json:"item_id" gorm:"not null"`
	TransactionID int       `json:"transaction_id"`
	Price         int       `json:"price"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (Purchase) TableName() string {
	return "purchases"
}
