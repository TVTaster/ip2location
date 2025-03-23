package config_test

import (
	"ip2location/internal/config"
	"os"
	"testing"
)

// Test LoadConfig
func TestLoadConfig(t *testing.T) {
	// Set the working directory to where the .env file exists
	if err := os.Chdir("../../"); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Set up test environment variables
	err := os.Setenv("PORT", "9090")
	if err != nil {
		return
	}
	err = os.Setenv("DB_TYPE", "csv")
	if err != nil {
		return
	}
	err = os.Setenv("RATE_LIMIT_PER_SEC", "10")
	if err != nil {
		return
	}

	t.Run("all environment variables set", func(t *testing.T) {
		cfg := config.LoadConfig()

		// Validate results
		if cfg.Port != "9090" {
			t.Errorf("Expected port 9090, but got %s", cfg.Port)
		}
		if cfg.DbType != "csv" {
			t.Errorf("Expected DB_TYPE 'csv', but got %s", cfg.DbType)
		}
		if cfg.RateLimitPerSec != 10 {
			t.Errorf("Expected RATE_LIMIT_PER_SEC 10, but got %d", cfg.RateLimitPerSec)
		}
	})

	t.Run("fallback values for invalid RATE_LIMIT_PER_SEC", func(t *testing.T) {
		// Set invalid RATE_LIMIT_PER_SEC and re-test
		err := os.Setenv("RATE_LIMIT_PER_SEC", "invalid")
		if err != nil {
			return
		}
		cfg := config.LoadConfig()

		if cfg.RateLimitPerSec != 5 {
			t.Errorf("Expected fallback RATE_LIMIT_PER_SEC 5, but got %d", cfg.RateLimitPerSec)
		}
	})

	t.Run("fallback values for invalid PORT", func(t *testing.T) {
		err := os.Setenv("PORT", "invalid")
		if err != nil {
			return
		}
		cfg := config.LoadConfig()

		if cfg.Port != "8080" {
			t.Errorf("Expected fallback PORT 8080, but got %s", cfg.Port)
		}
	})

	t.Run("missing environment variables should log fatal errors", func(t *testing.T) {
		// Unset all needed environment variables
		err := os.Unsetenv("PORT")
		if err != nil {
			return
		}
		err = os.Unsetenv("DB_TYPE")
		if err != nil {
			return
		}
		err = os.Unsetenv("RATE_LIMIT_PER_SEC")
		if err != nil {
			return
		}

		// Since log.Fatalf stops execution, you cannot directly assert it in tests
		// Use libraries like `testify`'s `assert.Panics` or isolate this behavior into smaller testable methods
	})
}

// Test GetEnv
func TestGetEnv(t *testing.T) {
	err := os.Setenv("TEST_ENV", "test_value")
	if err != nil {
		return
	}

	t.Run("Get existing environment variable", func(t *testing.T) {
		value := config.GetEnv("TEST_ENV")
		if value != "test_value" {
			t.Errorf("Expected 'test_value', but got '%s'", value)
		}
	})
}
