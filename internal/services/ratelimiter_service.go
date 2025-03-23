package services

import (
	"ip2location/internal/models"
	"time"
)

func NewRateLimiter(limit int) *models.RateLimiter {

	if limit <= 0 {
		limit = 5
	}

	return &models.RateLimiter{
		LimitPerSec:   limit,
		LastResetTime: time.Now(),
	}
}
