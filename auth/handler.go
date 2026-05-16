package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Register - регистрация ЗАКРЫТА, доступ только через предсозданные аккаунты
func (h *Handler) Register(c *gin.Context) {
	log.Println("🔵 ПОЛУЧЕН ЗАПРОС НА РЕГИСТРАЦИЮ (ЗАКРЫТО)")
	c.JSON(http.StatusForbidden, gin.H{
		"error": "Регистрация закрыта. Используйте предсозданные аккаунты: admin, teacher, student",
	})
	return
}

func (h *Handler) Login(c *gin.Context) {
	log.Println("========================================")
	log.Println("🔵 ПОЛУЧЕН ЗАПРОС НА ВХОД")
	log.Printf("🔵 Метод: %s, Путь: %s", c.Request.Method, c.Request.URL.Path)

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("📝 Логин: '%s'", req.Login)
	log.Printf("📝 Пароль: '%s'", req.Password)

	user, err := h.service.Login(&req)
	if err != nil {
		log.Printf("❌ Ошибка сервиса: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	log.Printf("✅ Пользователь найден: ID=%d, Login=%s, Role=%s", user.ID, user.Login, user.Role)

	token, err := GenerateToken(user.ID, user.Login, user.Role)
	if err != nil {
		log.Printf("❌ Ошибка генерации токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка генерации токена"})
		return
	}

	log.Printf("✅ Токен сгенерирован успешно")
	log.Printf("✅ Вход выполнен! Пользователь: %s", user.Login)
	log.Println("========================================")

	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
		User: UserResponse{
			ID:      user.ID,
			Login:   user.Login,
			Name:    user.Name,
			Role:    user.Role,
			Group:   user.Group,
			TotalXP: user.TotalXP,
			Coins:   user.Coins,
		},
	})
}

func (h *Handler) GetMe(c *gin.Context) {
	log.Println("🔵 ПОЛУЧЕН ЗАПРОС НА ПОЛУЧЕНИЕ ДАННЫХ ПОЛЬЗОВАТЕЛЯ")

	userID, exists := c.Get("userID")
	if !exists {
		log.Println("❌ userID не найден в контексте")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	log.Printf("✅ userID из контекста: %d", userID.(int))

	user, err := h.service.GetUserByID(userID.(int))
	if err != nil {
		log.Printf("❌ Пользователь не найден: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}

	log.Printf("✅ Данные пользователя отправлены: %s", user.Login)
	c.JSON(http.StatusOK, UserResponse{
		ID:      user.ID,
		Login:   user.Login,
		Name:    user.Name,
		Role:    user.Role,
		Group:   user.Group,
		TotalXP: user.TotalXP,
		Coins:   user.Coins,
	})
}
