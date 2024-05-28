package ratelimiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type RateLimiter struct {
    limiter *rate.Limiter
}

func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
    return &RateLimiter{
        limiter: rate.NewLimiter(r, b),
    }
}

func (rl *RateLimiter) Limit() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !rl.limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
            return
        }
        c.Next()
    }
}