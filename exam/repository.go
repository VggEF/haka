package exam

import (
	"database/sql"
	"fmt"
	"strings"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Вопросы ==========
func (r *Repository) CreateQuestion(q *models.ExamQuestion) error {
	query := `
        INSERT INTO exam_questions (question, answer, difficulty, subject, category, created_by)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, q.Question, q.Answer, q.Difficulty,
		q.Subject, q.Category, q.CreatedBy).Scan(&q.ID, &q.CreatedAt)
}

func (r *Repository) UpdateQuestion(id int, updates map[string]interface{}) error {
	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE exam_questions SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) DeleteQuestion(id int) error {
	query := `DELETE FROM exam_questions WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindQuestionByID(id int) (*models.ExamQuestion, error) {
	var q models.ExamQuestion
	query := `
        SELECT id, question, answer, difficulty, subject, category, created_by, created_at
        FROM exam_questions
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&q.ID, &q.Question, &q.Answer,
		&q.Difficulty, &q.Subject, &q.Category, &q.CreatedBy, &q.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *Repository) GetAllQuestions(subject, difficulty string, limit, offset int) ([]models.ExamQuestion, error) {
	query := `
        SELECT id, question, answer, difficulty, subject, category, created_at
        FROM exam_questions
        WHERE 1=1
    `
	args := []interface{}{}
	i := 1

	if subject != "" {
		query += fmt.Sprintf(" AND subject = $%d", i)
		args = append(args, subject)
		i++
	}
	if difficulty != "" {
		query += fmt.Sprintf(" AND difficulty = $%d", i)
		args = append(args, difficulty)
		i++
	}

	if limit == 0 {
		limit = 50
	}
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.ExamQuestion
	for rows.Next() {
		var q models.ExamQuestion
		err := rows.Scan(&q.ID, &q.Question, &q.Answer, &q.Difficulty,
			&q.Subject, &q.Category, &q.CreatedAt)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (r *Repository) GetQuestionsBySubject(subject string, difficulty string, limit int) ([]models.ExamQuestion, error) {
	query := `
        SELECT id, question, answer, difficulty, subject, category
        FROM exam_questions
        WHERE subject = $1
    `
	args := []interface{}{subject}
	i := 2

	if difficulty != "" {
		query += fmt.Sprintf(" AND difficulty = $%d", i)
		args = append(args, difficulty)
		i++
	}

	query += fmt.Sprintf(" ORDER BY RANDOM() LIMIT $%d", i)
	args = append(args, limit)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.ExamQuestion
	for rows.Next() {
		var q models.ExamQuestion
		err := rows.Scan(&q.ID, &q.Question, &q.Answer, &q.Difficulty,
			&q.Subject, &q.Category)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

// ========== Результаты экзаменов ==========
func (r *Repository) SaveExamResult(result *models.ExamResult) error {
	query := `
        INSERT INTO exam_results (user_id, subject, score, total_questions, correct_answers)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	return r.db.QueryRow(query, result.UserID, result.Subject, result.Score,
		result.TotalQuestions, result.CorrectAnswers).Scan(&result.ID)
}

func (r *Repository) GetUserExamResults(userID int, subject string, limit, offset int) ([]models.ExamResult, error) {
	query := `
        SELECT id, user_id, subject, score, total_questions, correct_answers, taken_at
        FROM exam_results
        WHERE user_id = $1
    `
	args := []interface{}{userID}
	i := 2

	if subject != "" {
		query += fmt.Sprintf(" AND subject = $%d", i)
		args = append(args, subject)
		i++
	}

	if limit == 0 {
		limit = 20
	}
	query += fmt.Sprintf(" ORDER BY taken_at DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.ExamResult
	for rows.Next() {
		var r models.ExamResult
		err := rows.Scan(&r.ID, &r.UserID, &r.Subject, &r.Score,
			&r.TotalQuestions, &r.CorrectAnswers, &r.TakenAt)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}

func (r *Repository) GetBestScore(userID int, subject string) (int, error) {
	var bestScore int
	query := `
        SELECT COALESCE(MAX(score), 0)
        FROM exam_results
        WHERE user_id = $1 AND subject = $2
    `
	err := r.db.QueryRow(query, userID, subject).Scan(&bestScore)
	return bestScore, err
}
