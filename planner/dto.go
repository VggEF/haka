package planner

import "time"

type CreateHabitRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	XPReward    int    `json:"xp_reward"`
}

type UpdateHabitRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	XPReward    int    `json:"xp_reward"`
	IsActive    bool   `json:"is_active"`
}

type HabitResponse struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	XPReward      int        `json:"xp_reward"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	Streak        int        `json:"streak"`
	IsCompleted   bool       `json:"is_completed"`
	LastCompleted *time.Time `json:"last_completed,omitempty"`
}

type CompleteHabitRequest struct {
	HabitID int `json:"habit_id" binding:"required"`
}

type HabitLogResponse struct {
	ID            int       `json:"id"`
	HabitID       int       `json:"habit_id"`
	HabitName     string    `json:"habit_name"`
	CompletedDate time.Time `json:"completed_date"`
	XPEarned      int       `json:"xp_earned"`
}

type GetHabitsQuery struct {
	IsActive    bool `form:"is_active"`
	IsCompleted bool `form:"is_completed"`
}
