package models

import "time"

type ExamQuestion struct {
	ID         int       `json:"id"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	Difficulty string    `json:"difficulty"`
	Subject    string    `json:"subject"`
	Category   string    `json:"category"`
	CreatedBy  int       `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type ExamResult struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	Subject        string    `json:"subject"`
	Score          int       `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	CorrectAnswers int       `json:"correct_answers"`
	TakenAt        time.Time `json:"taken_at"`
}
