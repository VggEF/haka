package planner

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

// ========== Пользовательские эндпоинты ==========

// GET /api/planner/habits
func (h *Handler) GetUserHabits(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	habits, err := h.service.GetUserHabits(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habits)
}

// POST /api/planner/habits/complete
func (h *Handler) CompleteHabit(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	var req CompleteHabitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CompleteHabit(userID.(int), req.HabitID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "привычка выполнена"})
}

// GET /api/planner/logs
func (h *Handler) GetHabitLogs(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	habitID, _ := strconv.Atoi(c.DefaultQuery("habit_id", "0"))

	logs, err := h.service.GetHabitLogs(userID.(int), habitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

// GET /api/planner/stats/weekly
func (h *Handler) GetWeeklyStats(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	stats, err := h.service.GetWeeklyStats(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// ========== Админские эндпоинты ==========

// GET /api/admin/planner/habits
func (h *Handler) GetAllHabits(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	isActive := c.DefaultQuery("is_active", "true") == "true"

	habits, err := h.service.GetAllHabits(isActive)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, habits)
}

// POST /api/admin/planner/habits
func (h *Handler) CreateHabit(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	userID, _ := c.Get("userID")

	var req CreateHabitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit, err := h.service.CreateHabit(&req, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, habit)
}

// PUT /api/admin/planner/habits/:id
func (h *Handler) UpdateHabit(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	var req UpdateHabitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateHabit(id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "привычка обновлена"})
}

// DELETE /api/admin/planner/habits/:id
func (h *Handler) DeleteHabit(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	if err := h.service.DeleteHabit(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "привычка удалена"})
}
