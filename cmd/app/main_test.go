// package services_test
package services_test

import (
	"errors"
	"ip2location/internal/mocks"
	"ip2location/internal/models"
	"ip2location/internal/services"
	"testing"
)

func TestServer_GetLocationByIP(t *testing.T) {
	// Mock response for IP lookup
	mockResponse := &models.IPInfo{
		Country: "United States",
		City:    "New York",
	}

	// Initialize the mock datastore
	mockDataStore := &mocks.MockDataStore{
		MockGetLocationByIP: func(ip string) (*models.IPInfo, error) {
			if ip == "8.8.8.8" {
				return mockResponse, nil
			}
			if ip == "invalid" {
				return nil, errors.New("invalid IP address")
			}
			return nil, errors.New("location not found")
		},
	}

	// Initialize the server with the mock datastore
	server := services.NewServer("8080", mockDataStore, nil)

	t.Run("valid IP address", func(t *testing.T) {
		location, err := server.Datastore.GetLocationByIP("8.8.8.8")
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if location.Country != "United States" || location.City != "New York" {
			t.Errorf("Unexpected location: %+v", location)
		}
	})

	t.Run("invalid IP address", func(t *testing.T) {
		_, err := server.Datastore.GetLocationByIP("invalid")
		if err == nil {
			t.Fatal("Expected an error for invalid IP, got nil")
		}
		if err.Error() != "invalid IP address" {
			t.Errorf("Expected error 'invalid IP address', got: %s", err)
		}
	})

	t.Run("location not found", func(t *testing.T) {
		_, err := server.Datastore.GetLocationByIP("1.1.1.1")
		if err == nil {
			t.Fatal("Expected an error for unknown IP, got nil")
		}
		if err.Error() != "location not found" {
			t.Errorf("Expected error 'location not found', got: %s", err)
		}
	})
}
