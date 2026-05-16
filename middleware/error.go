package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Логируем панику
				debug.PrintStack()
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Внутренняя ошибка сервера",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Ресурс не найден",
		})
	}
}
