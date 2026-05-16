package models

import "time"

type Schedule struct {
	ID         int       `json:"id"`
	GroupName  string    `json:"group_name"`
	GroupID    int       `json:"group_id"`
	Discipline string    `json:"discipline"`
	Teacher    string    `json:"teacher"`
	TeacherID  int       `json:"teacher_id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
	Date       time.Time `json:"date"`
	DayOfWeek  int       `json:"day_of_week"`
	Audience   string    `json:"audience"`
	LessonType string    `json:"lesson_type"`
	WeekType   int       `json:"week_type"`
	Comment    string    `json:"comment"`
	IsFromAPI  bool      `json:"is_from_api"`
	CreatedBy  int       `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type ScheduleCache struct {
	ID        int       `json:"id"`
	GroupID   int       `json:"group_id"`
	GroupName string    `json:"group_name"`
	Data      string    `json:"data"` // JSONB
	FetchedAt time.Time `json:"fetched_at"`
}
