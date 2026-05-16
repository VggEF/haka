package coins

import "time"

type AddCoinsRequest struct {
	UserID int    `json:"user_id" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
	Reason string `json:"reason"`
}

type SpendCoinsRequest struct {
	Amount int    `json:"amount" binding:"required"`
	Reason string `json:"reason"`
}

type BuyItemRequest struct {
	ItemID int `json:"item_id" binding:"required"`
}

type TransactionResponse struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	Type      string    `json:"type"` // earn, spend
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}

type ShopItemResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Icon        string `json:"icon"`
	Category    string `json:"category"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}

type GetTransactionsQuery struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
