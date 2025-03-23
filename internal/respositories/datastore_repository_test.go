package repositories_test

import (
	"ip2location/internal/models"
	"ip2location/internal/respositories"
	"os"
	"testing"
)

func TestInitDataStoreRepository(t *testing.T) {

	// Set the working directory to where the path to db is correct file exists
	if err := os.Chdir("../../"); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	t.Run("valid database type", func(t *testing.T) {

		// Example database type: "csv"
		dbType := "csv"
		err := os.Setenv("CSV_DATA_SOURCE_PATH", "resources/csv/test.csv")
		if err != nil {
			return
		}

		repo, err := repositories.InitDataStoreRepository(dbType)
		if err != nil {
			t.Fatalf("Expected no error for valid dbType '%s', got: %v", dbType, err)
		}

		if repo == nil {
			t.Fatalf("Expected a valid repository instance, got nil")
		}

		// Optional: Check if repo adheres to the `models.DataStore` interface
		_, ok := repo.(models.DataStore)
		if !ok {
			t.Errorf("Expected repo to implement models.DataStore")
		}
	})

	t.Run("invalid database type", func(t *testing.T) {
		dbType := "invalid_db"

		repo, err := repositories.InitDataStoreRepository(dbType)
		if repo != nil {
			t.Errorf("Expected nil repository for invalid dbType '%s', got: %v", dbType, repo)
		}

		if err == nil {
			t.Fatalf("Expected an error for invalid dbType '%s', got none", dbType)
		}

		// Check specific error message (optional)
		expectedError := "unsupported DB_TYPE: invalid_db"
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', got: '%s'", expectedError, err.Error())
		}
	})
}

func TestDataStoreMethods(t *testing.T) {
	// Mock implementation for testing specific methods
	mockDataStore := &MockDataStore{}

	t.Run("example datastore method", func(t *testing.T) {
		// Example: Mock query execution (replace with actual method from your DataStore interface)
		result, err := mockDataStore.SomeMethod()
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Validate result
		expected := "some expected result"
		if result != expected {
			t.Errorf("Expected result: '%s', got: '%s'", expected, result)
		}
	})
}

// Mock implementation of models.DataStore as an example
type MockDataStore struct{}

func (m *MockDataStore) SomeMethod() (string, error) {
	// Example implementation
	return "some expected result", nil
}
