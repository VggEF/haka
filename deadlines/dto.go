package deadlines

import "time"

type CreateDeadlineRequest struct {
	Title       string `json:"title" binding:"required"`
	Subject     string `json:"subject" binding:"required"`
	GroupName   string `json:"group_name"`
	DueDate     string `json:"due_date" binding:"required"`
	DueTime     string `json:"due_time"`
	Priority    string `json:"priority"` // high, medium, low
	Description string `json:"description"`
	AssignedTo  int    `json:"assigned_to"`
}

type UpdateDeadlineRequest struct {
	Title       string `json:"title"`
	Subject     string `json:"subject"`
	GroupName   string `json:"group_name"`
	DueDate     string `json:"due_date"`
	DueTime     string `json:"due_time"`
	Priority    string `json:"priority"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type DeadlineResponse struct {
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
	CreatedAt   time.Time `json:"created_at"`
	DaysLeft    int       `json:"days_left"`
	IsOverdue   bool      `json:"is_overdue"`
}

type GetDeadlinesQuery struct {
	GroupName string `form:"group_name"`
	Subject   string `form:"subject"`
	Status    string `form:"status"` // pending, completed, overdue
	Priority  string `form:"priority"`
	DueFrom   string `form:"due_from"`
	DueTo     string `form:"due_to"`
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
}

type CompleteDeadlineRequest struct {
	Status string `json:"status"`
}
