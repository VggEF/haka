package checklist

import "time"

type CreateChecklistItemRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Category    string `json:"category"`
	XP          int    `json:"xp"`
	SortOrder   int    `json:"sort_order"`
}

type UpdateChecklistItemRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	XP          int    `json:"xp"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active"`
}

type ChecklistItemResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	XP          int        `json:"xp"`
	SortOrder   int        `json:"sort_order"`
	IsActive    bool       `json:"is_active"`
	IsCompleted bool       `json:"is_completed"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type CompleteItemRequest struct {
	ItemID int `json:"item_id" binding:"required"`
}

type GetChecklistQuery struct {
	Category  string `form:"category"`
	Completed bool   `form:"completed"`
}
