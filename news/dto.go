package news

import "time"

type CreateNewsRequest struct {
	Title     string `json:"title" binding:"required"`
	ShortText string `json:"short_text" binding:"required"`
	FullText  string `json:"full_text"`
	ImageURL  string `json:"image_url"`
	Category  string `json:"category"`
	IsPinned  bool   `json:"is_pinned"`
}

type UpdateNewsRequest struct {
	Title     string `json:"title"`
	ShortText string `json:"short_text"`
	FullText  string `json:"full_text"`
	ImageURL  string `json:"image_url"`
	Category  string `json:"category"`
	IsPinned  bool   `json:"is_pinned"`
}

type NewsResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ShortText string    `json:"short_text"`
	FullText  string    `json:"full_text"`
	ImageURL  string    `json:"image_url"`
	Category  string    `json:"category"`
	Date      time.Time `json:"date"`
	IsPinned  bool      `json:"is_pinned"`
	Views     int       `json:"views"`
	CreatedAt time.Time `json:"created_at"`
}

type GetNewsQuery struct {
	Category string `form:"category"`
	Pinned   bool   `form:"pinned"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}
