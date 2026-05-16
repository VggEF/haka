package models

import "time"

type Deadline struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Subject     string    `json:"subject"`
	GroupName   string    `json:"group_name"`
	DueDate     time.Time `json:"due_date"`
	DueTime     string    `json:"due_time"`
	Priority    string    `json:"priority"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedBy   int       `json:"created_by"`
	AssignedTo  int       `json:"assigned_to"`
	CreatedAt   time.Time `json:"created_at"`
}
