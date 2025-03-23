package services_test

import (
	"ip2location/internal/services"
	"testing"
	"time"
)

func TestNewRateLimiter(t *testing.T) {
	// Test case: Valid rate limiter initialization
	t.Run("valid initialization", func(t *testing.T) {
		limit := 5 // e.g., 5 requests per second
		rateLimiter := services.NewRateLimiter(limit)

		if rateLimiter == nil {
			t.Fatalf("NewRateLimiter returned nil")
		}

		if rateLimiter.LimitPerSec != limit {
			t.Errorf("Expected limit: %d, got: %d", limit, rateLimiter.LimitPerSec)
		}
	})

	// Test case: Zero or negative limit
	t.Run("invalid initialization", func(t *testing.T) {
		invalidLimits := []int{0, -1}

		for _, limit := range invalidLimits {
			rateLimiter := services.NewRateLimiter(limit)
			if rateLimiter == nil {
				t.Fatalf("NewRateLimiter should not return nil for limit %d", limit)
			}
			if rateLimiter.LimitPerSec != 5 {
				t.Errorf("Expected limit to default to 5, got: %d", rateLimiter.LimitPerSec)
			}
		}
	})
}

func TestRateLimiterFunctionality(t *testing.T) {
	limit := 2 // 2 requests per second
	rateLimiter := services.NewRateLimiter(limit)
	if rateLimiter == nil {
		t.Fatalf("NewRateLimiter returned nil")
	}

	t.Run("allows up to limit", func(t *testing.T) {
		for i := 0; i < limit; i++ {
			if !rateLimiter.Allow() {
				t.Errorf("Request %d/%d was not allowed by rate limiter", i+1, limit)
			}
		}

		// Next request should fail since limit is reached.
		if rateLimiter.Allow() {
			t.Errorf("Exceeding request limit was allowed")
		}
	})

	t.Run("resets after time period", func(t *testing.T) {
		time.Sleep(1 * time.Second) // Adjust to match implementation timer
		if !rateLimiter.Allow() {
			t.Errorf("Rate limiter did not reset after time period")
		}
	})
}
