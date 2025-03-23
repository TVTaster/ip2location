package mocks

import (
	"errors"
	"ip2location/internal/models"
)

type MockDataStore struct {
	// You can use function fields to allow dynamic behavior
	MockGetLocationByIP func(ip string) (*models.IPInfo, error)
}

func (m *MockDataStore) GetLocationByIP(ip string) (*models.IPInfo, error) {
	// Delegate to the mock function
	if m.MockGetLocationByIP != nil {
		return m.MockGetLocationByIP(ip)
	}
	// Default behavior (optional)
	//	return models.IPInfo{Country: "United States", City: "New York"}, nil
	return nil, nil
}

// MockDataStoreRepository implements the DataStoreRepository interface
type MockDataStoreRepository struct {
	MockData  map[string]models.IPInfo // Simulated IP to IPInfo mapping
	lastQuery string
	forceErr  error
}

// Query simulates the behavior of fetching location data for a given IP
func (m *MockDataStoreRepository) Query(ip string) (models.IPInfo, error) {
	m.lastQuery = ip
	if m.forceErr != nil {
		return models.IPInfo{}, m.forceErr
	}

	// Simulate a query result
	if result, exists := m.MockData[ip]; exists {
		return result, nil
	}
	return models.IPInfo{}, errors.New("not found")
}
