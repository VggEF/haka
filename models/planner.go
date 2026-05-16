package models

import "time"

type Habit struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	XPReward    int       `json:"xp_reward"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserHabit struct {
	UserID        int        `json:"user_id"`
	HabitID       int        `json:"habit_id"`
	Streak        int        `json:"streak"`
	LastCompleted *time.Time `json:"last_completed"`
}

type HabitLog struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	HabitID       int       `json:"habit_id"`
	CompletedDate time.Time `json:"completed_date"`
	XPEarned      int       `json:"xp_earned"`
}
