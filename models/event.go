package models

import "time"

type Event struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	ShortText      string    `json:"short_text"`
	FullText       string    `json:"full_text"`
	Date           time.Time `json:"date"`
	Time           string    `json:"time"`
	Type           string    `json:"type"`
	Category       string    `json:"category"`
	Location       string    `json:"location"`
	Price          string    `json:"price"`
	Organizer      string    `json:"organizer"`
	ImageURL       string    `json:"image_url"`
	AvailableSpots int       `json:"available_spots"`
	Contact        string    `json:"contact"`
	Registrations  string    `json:"registrations"` // JSONB поле
	CreatedBy      int       `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
}

type EventRegistration struct {
	ID           int       `json:"id"`
	EventID      int       `json:"event_id"`
	UserID       int       `json:"user_id"`
	RegisteredAt time.Time `json:"registered_at"`
	Status       string    `json:"status"`
}
