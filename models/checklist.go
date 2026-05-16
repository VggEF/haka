package models

import "time"

type ChecklistItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	XP          int    `json:"xp"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active"`
}

type UserChecklist struct {
	UserID      int        `json:"user_id"`
	ItemID      int        `json:"item_id"`
	CompletedAt *time.Time `json:"completed_at"`
	IsCompleted bool       `json:"is_completed"`
}
