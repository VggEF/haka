package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireRole создает middleware для проверки роли
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := GetUserRole(c)
		if userRole != role {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Доступ запрещен. Требуется роль: " + role,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireAnyRole создает middleware для проверки наличия любой из ролей
func RequireAnyRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := GetUserRole(c)
		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Доступ запрещен. Требуется одна из ролей: " + strings.Join(roles, ", "),
		})
		c.Abort()
	}
}

// RequireSelfOrAdmin проверяет, что пользователь получает свои данные или является админом
func RequireSelfOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := GetUserID(c)
		userRole := GetUserRole(c)

		// Админам разрешено всё
		if userRole == "admin" {
			c.Next()
			return
		}

		// Получаем ID из параметра
		paramID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			c.Abort()
			return
		}

		// Проверяем, что запрашивает свои данные
		if userID != paramID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
			c.Abort()
			return
		}

		c.Next()
	}
}
