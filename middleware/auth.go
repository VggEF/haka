package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Используем тот же секрет, что и в auth пакете
// Лучше брать из конфига
var jwtSecret = []byte("your-secret-key-change-in-production")

type Claims struct {
	UserID int    `json:"user_id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// AuthMiddleware проверяет JWT токен
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
			c.Abort()
			return
		}

		// Проверяем формат Bearer
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат авторизации. Используйте: Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Проверяем метод подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен: " + err.Error()})
			c.Abort()
			return
		}

		// Сохраняем в контекст (используем camelCase для консистентности)
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Set("userLogin", claims.Login)

		c.Next()
	}
}

// RoleMiddleware проверяет, имеет ли пользователь одну из разрешенных ролей
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка определения роли"})
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if roleStr == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "Доступ запрещен. Требуется роль: " + strings.Join(allowedRoles, ", "),
		})
		c.Abort()
	}
}

// AdminOnlyMiddleware только для админов
func AdminOnlyMiddleware() gin.HandlerFunc {
	return RoleMiddleware("admin")
}

// TeacherOrAdminMiddleware для преподавателей и админов
func TeacherOrAdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware("admin", "teacher")
}

// StudentMiddleware только для студентов
func StudentMiddleware() gin.HandlerFunc {
	return RoleMiddleware("student")
}

// GetUserRole возвращает роль текущего пользователя из контекста
func GetUserRole(c *gin.Context) string {
	role, exists := c.Get("userRole")
	if !exists {
		return ""
	}
	return role.(string)
}

// GetUserID возвращает ID текущего пользователя из контекста
func GetUserID(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(int)
}

// CheckRole проверяет, соответствует ли роль пользователя ожидаемой
func CheckRole(c *gin.Context, expectedRole string) bool {
	userRole := GetUserRole(c)
	return userRole == expectedRole
}
