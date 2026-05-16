package exam

import "time"

type CreateQuestionRequest struct {
	Question   string `json:"question" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
	Difficulty string `json:"difficulty"` // easy, medium, hard
	Subject    string `json:"subject" binding:"required"`
	Category   string `json:"category"`
}

type UpdateQuestionRequest struct {
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Difficulty string `json:"difficulty"`
	Subject    string `json:"subject"`
	Category   string `json:"category"`
}

type QuestionResponse struct {
	ID         int       `json:"id"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	Difficulty string    `json:"difficulty"`
	Subject    string    `json:"subject"`
	Category   string    `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
}

type GenerateExamRequest struct {
	Subject        string `json:"subject" binding:"required"`
	Difficulty     string `json:"difficulty"`
	QuestionsCount int    `json:"questions_count"`
}

type ExamQuestionResponse struct {
	ID         int    `json:"id"`
	Question   string `json:"question"`
	Difficulty string `json:"difficulty"`
}

type SubmitExamRequest struct {
	Subject string         `json:"subject" binding:"required"`
	Answers map[int]string `json:"answers" binding:"required"`
}

type ExamResultResponse struct {
	Score          int              `json:"score"`
	TotalQuestions int              `json:"total_questions"`
	CorrectAnswers int              `json:"correct_answers"`
	Percentage     float64          `json:"percentage"`
	Grade          string           `json:"grade"`
	Results        []QuestionResult `json:"results"`
}

type QuestionResult struct {
	QuestionID    int    `json:"question_id"`
	Question      string `json:"question"`
	UserAnswer    string `json:"user_answer"`
	CorrectAnswer string `json:"correct_answer"`
	IsCorrect     bool   `json:"is_correct"`
}

type GetQuestionsQuery struct {
	Subject    string `form:"subject"`
	Difficulty string `form:"difficulty"`
	Limit      int    `form:"limit"`
	Offset     int    `form:"offset"`
}
