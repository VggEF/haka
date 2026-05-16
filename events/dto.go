package events

import "time"

type CreateEventRequest struct {
	Title          string `json:"title" binding:"required"`
	ShortText      string `json:"short_text" binding:"required"`
	FullText       string `json:"full_text"`
	Date           string `json:"date" binding:"required"`
	Time           string `json:"time"`
	Type           string `json:"type" binding:"required"` // sport, school, student
	Category       string `json:"category"`
	Location       string `json:"location"`
	Price          string `json:"price"`
	Organizer      string `json:"organizer"`
	ImageURL       string `json:"image_url"`
	AvailableSpots int    `json:"available_spots"`
	Contact        string `json:"contact"`
}

type UpdateEventRequest struct {
	Title          string `json:"title"`
	ShortText      string `json:"short_text"`
	FullText       string `json:"full_text"`
	Date           string `json:"date"`
	Time           string `json:"time"`
	Type           string `json:"type"`
	Category       string `json:"category"`
	Location       string `json:"location"`
	Price          string `json:"price"`
	Organizer      string `json:"organizer"`
	ImageURL       string `json:"image_url"`
	AvailableSpots int    `json:"available_spots"`
	Contact        string `json:"contact"`
}

type EventResponse struct {
	ID                 int       `json:"id"`
	Title              string    `json:"title"`
	ShortText          string    `json:"short_text"`
	FullText           string    `json:"full_text"`
	Date               time.Time `json:"date"`
	Time               string    `json:"time"`
	Type               string    `json:"type"`
	Category           string    `json:"category"`
	Location           string    `json:"location"`
	Price              string    `json:"price"`
	Organizer          string    `json:"organizer"`
	ImageURL           string    `json:"image_url"`
	AvailableSpots     int       `json:"available_spots"`
	Contact            string    `json:"contact"`
	RegistrationsCount int       `json:"registrations_count"`
	IsRegistered       bool      `json:"is_registered"`
}

type RegisterEventRequest struct {
	UserID int `json:"user_id"`
}

type GetEventsQuery struct {
	Type     string `form:"type"`
	Category string `form:"category"`
	DateFrom string `form:"date_from"`
	DateTo   string `form:"date_to"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}
