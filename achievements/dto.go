package achievements

import "time"

type CreateAchievementRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	XP          int    `json:"xp" binding:"required"`
	Icon        string `json:"icon"`
	Rarity      string `json:"rarity"` // common, rare, epic, legendary
	Category    string `json:"category"`
	Condition   string `json:"condition"`
}

type UpdateAchievementRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	XP          int    `json:"xp"`
	Icon        string `json:"icon"`
	Rarity      string `json:"rarity"`
	Category    string `json:"category"`
	Condition   string `json:"condition"`
}

type AchievementResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	XP          int        `json:"xp"`
	Icon        string     `json:"icon"`
	Rarity      string     `json:"rarity"`
	Category    string     `json:"category"`
	Condition   string     `json:"condition"`
	CreatedAt   time.Time  `json:"created_at"`
	IsUnlocked  bool       `json:"is_unlocked"`
	UnlockedAt  *time.Time `json:"unlocked_at,omitempty"`
}

type AwardAchievementRequest struct {
	UserID        int `json:"user_id" binding:"required"`
	AchievementID int `json:"achievement_id" binding:"required"`
}

type AwardToGroupRequest struct {
	GroupName     string `json:"group_name" binding:"required"`
	AchievementID int    `json:"achievement_id" binding:"required"`
}

type CreateDailyQuestRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	XP          int       `json:"xp" binding:"required"`
	Coins       int       `json:"coins"`
	Requirement string    `json:"requirement"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type DailyQuestResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	XP          int       `json:"xp"`
	Coins       int       `json:"coins"`
	Requirement string    `json:"requirement"`
	ExpiresAt   time.Time `json:"expires_at"`
	IsCompleted bool      `json:"is_completed"`
}

type GetUserAchievementsQuery struct {
	Category string `form:"category"`
	Unlocked bool   `form:"unlocked"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}
