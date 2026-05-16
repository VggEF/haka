package models

import "time"

type News struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ShortText string    `json:"short_text"`
	FullText  string    `json:"full_text"`
	ImageURL  string    `json:"image_url"`
	Category  string    `json:"category"`
	Date      time.Time `json:"date"`
	IsPinned  bool      `json:"is_pinned"`
	CreatedBy int       `json:"created_by"`
	Views     int       `json:"views"`
	CreatedAt time.Time `json:"created_at"`
}
