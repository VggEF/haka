package exam

import (
	"errors"
	"math/rand"
	"strings"
	"student-app/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// ========== Управление вопросами (админ/teacher) ==========
func (s *Service) CreateQuestion(req *CreateQuestionRequest, userID int) (*models.ExamQuestion, error) {
	difficulty := req.Difficulty
	if difficulty == "" {
		difficulty = "medium"
	}

	q := &models.ExamQuestion{
		Question:   req.Question,
		Answer:     req.Answer,
		Difficulty: difficulty,
		Subject:    req.Subject,
		Category:   req.Category,
		CreatedBy:  userID,
	}

	if err := s.repo.CreateQuestion(q); err != nil {
		return nil, err
	}
	return q, nil
}

func (s *Service) UpdateQuestion(id int, req *UpdateQuestionRequest) error {
	updates := make(map[string]interface{})

	if req.Question != "" {
		updates["question"] = req.Question
	}
	if req.Answer != "" {
		updates["answer"] = req.Answer
	}
	if req.Difficulty != "" {
		updates["difficulty"] = req.Difficulty
	}
	if req.Subject != "" {
		updates["subject"] = req.Subject
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}

	return s.repo.UpdateQuestion(id, updates)
}

func (s *Service) DeleteQuestion(id int) error {
	return s.repo.DeleteQuestion(id)
}

func (s *Service) GetQuestionByID(id int) (*QuestionResponse, error) {
	q, err := s.repo.FindQuestionByID(id)
	if err != nil {
		return nil, err
	}
	return &QuestionResponse{
		ID:         q.ID,
		Question:   q.Question,
		Answer:     q.Answer,
		Difficulty: q.Difficulty,
		Subject:    q.Subject,
		Category:   q.Category,
		CreatedAt:  q.CreatedAt,
	}, nil
}

func (s *Service) GetAllQuestions(query *GetQuestionsQuery) ([]QuestionResponse, int, error) {
	limit := query.Limit
	if limit == 0 {
		limit = 50
	}

	questions, err := s.repo.GetAllQuestions(query.Subject, query.Difficulty, limit, query.Offset)
	if err != nil {
		return nil, 0, err
	}

	var response []QuestionResponse
	for _, q := range questions {
		response = append(response, QuestionResponse{
			ID:         q.ID,
			Question:   q.Question,
			Answer:     q.Answer,
			Difficulty: q.Difficulty,
			Subject:    q.Subject,
			Category:   q.Category,
			CreatedAt:  q.CreatedAt,
		})
	}

	return response, len(questions), nil
}

// ========== Генерация экзамена ==========
func (s *Service) GenerateExam(req *GenerateExamRequest) ([]ExamQuestionResponse, error) {
	questionsCount := req.QuestionsCount
	if questionsCount == 0 {
		questionsCount = 10
	}

	questions, err := s.repo.GetQuestionsBySubject(req.Subject, req.Difficulty, questionsCount)
	if err != nil {
		return nil, err
	}

	if len(questions) == 0 {
		return nil, errors.New("нет вопросов по данному предмету")
	}

	// Перемешиваем вопросы
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	// Ограничиваем количество
	if len(questions) > questionsCount {
		questions = questions[:questionsCount]
	}

	var response []ExamQuestionResponse
	for _, q := range questions {
		response = append(response, ExamQuestionResponse{
			ID:         q.ID,
			Question:   q.Question,
			Difficulty: q.Difficulty,
		})
	}
	return response, nil
}

// ========== Проверка экзамена ==========
func (s *Service) SubmitExam(userID int, req *SubmitExamRequest) (*ExamResultResponse, error) {
	// Получаем все вопросы по ответам
	questionIDs := make([]int, 0, len(req.Answers))
	for id := range req.Answers {
		questionIDs = append(questionIDs, id)
	}

	var correctCount int
	var results []QuestionResult

	for questionID, userAnswer := range req.Answers {
		q, err := s.repo.FindQuestionByID(questionID)
		if err != nil {
			continue
		}

		isCorrect := normalizeString(userAnswer) == normalizeString(q.Answer)
		if isCorrect {
			correctCount++
		}

		results = append(results, QuestionResult{
			QuestionID:    q.ID,
			Question:      q.Question,
			UserAnswer:    userAnswer,
			CorrectAnswer: q.Answer,
			IsCorrect:     isCorrect,
		})
	}

	totalQuestions := len(req.Answers)
	percentage := float64(correctCount) / float64(totalQuestions) * 100

	// Определяем оценку
	grade := getGrade(percentage)

	// Сохраняем результат
	result := &models.ExamResult{
		UserID:         userID,
		Subject:        req.Subject,
		Score:          correctCount,
		TotalQuestions: totalQuestions,
		CorrectAnswers: correctCount,
	}
	s.repo.SaveExamResult(result)

	return &ExamResultResponse{
		Score:          correctCount,
		TotalQuestions: totalQuestions,
		CorrectAnswers: correctCount,
		Percentage:     percentage,
		Grade:          grade,
		Results:        results,
	}, nil
}

// ========== Статистика ==========
func (s *Service) GetUserResults(userID int, subject string) ([]ExamResultResponse, error) {
	results, err := s.repo.GetUserExamResults(userID, subject, 20, 0)
	if err != nil {
		return nil, err
	}

	var response []ExamResultResponse
	for _, r := range results {
		percentage := float64(r.CorrectAnswers) / float64(r.TotalQuestions) * 100
		response = append(response, ExamResultResponse{
			Score:          r.Score,
			TotalQuestions: r.TotalQuestions,
			CorrectAnswers: r.CorrectAnswers,
			Percentage:     percentage,
			Grade:          getGrade(percentage),
		})
	}
	return response, nil
}

func (s *Service) GetBestScore(userID int, subject string) (int, error) {
	return s.repo.GetBestScore(userID, subject)
}

// ========== Вспомогательные функции ==========
func normalizeString(s string) string {
	// Приводим к нижнему регистру и убираем лишние пробелы
	result := strings.ToLower(strings.TrimSpace(s))
	return result
}

func getGrade(percentage float64) string {
	switch {
	case percentage >= 90:
		return "5 (Отлично)"
	case percentage >= 75:
		return "4 (Хорошо)"
	case percentage >= 60:
		return "3 (Удовлетворительно)"
	default:
		return "2 (Неудовлетворительно)"
	}
}
