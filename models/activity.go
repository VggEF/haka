package models

import "time"

type UserActivity struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Action    string    `json:"action"`
	Page      string    `json:"page"`
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
}
