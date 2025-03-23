package services

import (
	"ip2location/internal/models"
)

func NewServer(port string, ds models.DataStore, rl *models.RateLimiter) *models.Server {
	return &models.Server{
		Port:        port,
		Datastore:   ds,
		RateLimiter: rl,
	}
}
