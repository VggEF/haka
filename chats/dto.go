package chats

import "time"

type CreateChatRequest struct {
	Name    string `json:"name"`
	Type    string `json:"type"` // private, group, course
	Members []int  `json:"members" binding:"required,min=1"`
}

type SendMessageRequest struct {
	Text    string `json:"text" binding:"required"`
	FileURL string `json:"file_url"`
}

type ChatResponse struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Type          string               `json:"type"`
	LastMessage   string               `json:"last_message"`
	LastMessageAt time.Time            `json:"last_message_at"`
	UnreadCount   int                  `json:"unread_count"`
	Members       []ChatMemberResponse `json:"members,omitempty"`
	CreatedAt     time.Time            `json:"created_at"`
}

type ChatMemberResponse struct {
	UserID   int       `json:"user_id"`
	UserName string    `json:"user_name"`
	Role     string    `json:"role"` // admin, member
	JoinedAt time.Time `json:"joined_at"`
}

type MessageResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	UserPhoto string    `json:"user_photo"`
	Text      string    `json:"text"`
	FileURL   string    `json:"file_url"`
	IsRead    bool      `json:"is_read"`
	IsMine    bool      `json:"is_mine"`
	CreatedAt time.Time `json:"created_at"`
}

type GetMessagesQuery struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type MarkAsReadRequest struct {
	MessageID int `json:"message_id" binding:"required"`
}
