package config

import (
	"github.com/joho/godotenv"
	"ip2location/internal/models"
	"log"
	"os"
	"strconv"
)

/*
	 LoadConfig loads the configuration for the application from environment variables
	 defined in a `.env` file. It relies on the `godotenv` package to parse the `.env` file
	 and retrieve the required environment variables.

	 The function retrieves the following *MUST HAVE* environment variables:
	 - PORT: The port number on which the application should run.
	 - DB_TYPE: The type of database being used.
	 - RATE_LIMIT_PER_SEC: The rate limit (in requests per second) for the application.

	 It converts the `RATE_LIMIT_PER_SEC` value to an integer, or falls back to a default
	 value of 5 if the conversion fails.

	 Returns:
	 - A pointer to a `models.Config` struct containing the loaded configuration.

	 Example:

		config := LoadConfig()
*/
func LoadConfig() *models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	port := GetEnv("PORT")
	dbType := GetEnv("DB_TYPE")
	rateLimitPerSec := GetEnv("RATE_LIMIT_PER_SEC")

	rateLimit, err := strconv.Atoi(rateLimitPerSec)
	if err != nil {
		log.Println("Atoi function failed, RATE_LIMIT_PER_SEC using fallback default value - 5")
		rateLimit = 5
	}

	_, err = strconv.Atoi(port)
	if err != nil {
		log.Println("Atoi function failed, PORT isn't a valid number, using fallback default value - 8080")
		port = "8080"
	}

	return &models.Config{
		Port:            port,
		RateLimitPerSec: rateLimit,
		DbType:          dbType,
	}
}

/*
	 GetEnv checks if the given environment variable is set and returns it.
	 If it is not set, the function logs a fatal error and stops the execution of the program.

	 Parameters:
	 - key: The name of the environment variable.

	 Example:

		GetEnv("PORT")
*/
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is not set", key)
	}
	return value
}
