package exam

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// ========== Студенческие эндпоинты ==========

// POST /api/exam/generate
func (h *Handler) GenerateExam(c *gin.Context) {
	var req GenerateExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questions, err := h.service.GenerateExam(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

// POST /api/exam/submit
func (h *Handler) SubmitExam(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	var req SubmitExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.SubmitExam(userID.(int), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET /api/exam/results
func (h *Handler) GetResults(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	subject := c.DefaultQuery("subject", "")

	results, err := h.service.GetUserResults(userID.(int), subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GET /api/exam/best/:subject
func (h *Handler) GetBestScore(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	subject := c.Param("subject")

	score, err := h.service.GetBestScore(userID.(int), subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"best_score": score})
}

// GET /api/exam/subjects
func (h *Handler) GetSubjects(c *gin.Context) {
	subjects := []string{"Программирование", "Математика", "Английский", "Физика", "Базы данных"}
	c.JSON(http.StatusOK, subjects)
}

// ========== Админские эндпоинты ==========

// GET /api/admin/exam/questions
func (h *Handler) GetAllQuestions(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" && userRole != "teacher" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	var query GetQuestionsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questions, total, err := h.service.GetAllQuestions(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  questions,
		"total": total,
	})
}

// GET /api/admin/exam/questions/:id
func (h *Handler) GetQuestionByID(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" && userRole != "teacher" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	question, err := h.service.GetQuestionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "вопрос не найден"})
		return
	}

	c.JSON(http.StatusOK, question)
}

// POST /api/admin/exam/questions
func (h *Handler) CreateQuestion(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" && userRole != "teacher" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	userID, _ := c.Get("userID")

	var req CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question, err := h.service.CreateQuestion(&req, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, question)
}

// PUT /api/admin/exam/questions/:id
func (h *Handler) UpdateQuestion(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" && userRole != "teacher" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	var req UpdateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateQuestion(id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "вопрос обновлен"})
}

// DELETE /api/admin/exam/questions/:id
func (h *Handler) DeleteQuestion(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" && userRole != "teacher" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	if err := h.service.DeleteQuestion(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "вопрос удален"})
}
