package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	visits map[string][]time.Time
	mu     sync.RWMutex
	limit  int
	window time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visits: make(map[string][]time.Time),
		limit:  limit,
		window: window,
	}

	// Запускаем горутину для очистки старых записей
	go rl.cleanup()

	return rl
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for ip, times := range rl.visits {
			var newTimes []time.Time
			for _, t := range times {
				if now.Sub(t) < rl.window {
					newTimes = append(newTimes, t)
				}
			}
			if len(newTimes) == 0 {
				delete(rl.visits, ip)
			} else {
				rl.visits[ip] = newTimes
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	times := rl.visits[ip]

	// Удаляем старые записи
	var validTimes []time.Time
	for _, t := range times {
		if now.Sub(t) < rl.window {
			validTimes = append(validTimes, t)
		}
	}

	if len(validTimes) >= rl.limit {
		return false
	}

	rl.visits[ip] = append(validTimes, now)
	return true
}

func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(limit, window)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !limiter.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Слишком много запросов. Попробуйте позже.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
