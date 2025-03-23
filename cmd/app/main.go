package services

import (
	"ip2location/internal/config"
	repositories "ip2location/internal/respositories"
	"ip2location/internal/services"
	"log"
)

func main() {

	log.Println("Loading configuration...")
	cfg := config.LoadConfig()

	log.Println("Initializing datastore...")
	datastore, err := repositories.InitDataStoreRepository(cfg.DbType)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating rate limiter...")
	rateLimiter := services.NewRateLimiter(cfg.RateLimitPerSec)

	log.Println("Creating server...")
	server := services.NewServer(cfg.Port, datastore, rateLimiter)

	server.Start()
}
