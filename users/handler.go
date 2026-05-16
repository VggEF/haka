package users

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

// GET /api/users/me
func (h *Handler) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	user, err := h.service.GetByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:       user.ID,
		Login:    user.Login,
		Name:     user.Name,
		Role:     user.Role,
		Group:    user.Group,
		Course:   user.Course,
		Faculty:  user.Faculty,
		Email:    user.Email,
		Phone:    user.Phone,
		Telegram: user.Telegram,
		VK:       user.VK,
		GitHub:   user.GitHub,
		Photo:    user.Photo,
		TotalXP:  user.TotalXP,
		Coins:    user.Coins,
	})
}

// PUT /api/users/me
func (h *Handler) UpdateMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(userID.(int), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "профиль обновлен"})
}

// GET /api/users/profile/student
func (h *Handler) GetStudentProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	profile, err := h.service.GetStudentProfile(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// PUT /api/users/profile/student
func (h *Handler) UpdateStudentProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	var req UpdateStudentProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStudentProfile(userID.(int), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "профиль обновлен"})
}

// GET /api/users/profile/teacher
func (h *Handler) GetTeacherProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	profile, err := h.service.GetTeacherProfile(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// PUT /api/users/profile/teacher
func (h *Handler) UpdateTeacherProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	var req UpdateTeacherProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateTeacherProfile(userID.(int), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "профиль обновлен"})
}

// GET /api/users (только для админа)
func (h *Handler) GetAll(c *gin.Context) {
	role := c.DefaultQuery("role", "student")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	users, err := h.service.GetAll(role, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []UserResponse
	for _, u := range users {
		response = append(response, UserResponse{
			ID:      u.ID,
			Login:   u.Login,
			Name:    u.Name,
			Role:    u.Role,
			Group:   u.Group,
			Course:  u.Course,
			Faculty: u.Faculty,
			Email:   u.Email,
			Phone:   u.Phone,
		})
	}

	c.JSON(http.StatusOK, response)
}

// DELETE /api/users/:id (только для админа)
func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "пользователь удален"})
}

// POST /api/users (только для админа)
func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:      user.ID,
		Login:   user.Login,
		Name:    user.Name,
		Role:    user.Role,
		Group:   user.Group,
		Course:  user.Course,
		Faculty: user.Faculty,
	})
}
