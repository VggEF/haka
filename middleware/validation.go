package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Неверный формат запроса: " + err.Error(),
			})
			c.Abort()
			return
		}

		if err := validate.Struct(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ошибка валидации: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("validated_body", obj)
		c.Next()
	}
}
