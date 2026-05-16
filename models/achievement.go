package models

import "time"

type Achievement struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	XP          int       `json:"xp"`
	Icon        string    `json:"icon"`
	Rarity      string    `json:"rarity"`
	Category    string    `json:"category"`
	Condition   string    `json:"condition"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserAchievement struct {
	UserID        int       `json:"user_id"`
	AchievementID int       `json:"achievement_id"`
	EarnedAt      time.Time `json:"earned_at"`
}

type DailyQuest struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	XP          int       `json:"xp"`
	Coins       int       `json:"coins"`
	Requirement string    `json:"requirement"`
	ExpiresAt   time.Time `json:"expires_at"`
}
