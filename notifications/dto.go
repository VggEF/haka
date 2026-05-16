package notifications

import "time"

type SendNotificationRequest struct {
	UserID  int    `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
	Type    string `json:"type"` // deadline, achievement, event, system
	Data    string `json:"data"`
}

type NotificationResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"is_read"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

type GetNotificationsQuery struct {
	Type   string `form:"type"`
	IsRead *bool  `form:"is_read"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}
