package repositories

import (
	"fmt"
	"ip2location/internal/config"
	"ip2location/internal/models"
)

type DataStoreRepository interface {
	GetLocationByIP(id string) (*models.IPInfo, error)
}

/*
InitDataStoreRepository initializes and returns a `DataStoreRepository` implementation
based on the provided `dbType` parameter.

Parameters:
  - dbType: A string indicating the type of data source repository to initialize.
    Supported values:
  - "csv": Initializes a repository that loads data from a CSV file.
  - "api": Initializes a repository that interacts with an API to fetch data.

If the `dbType` is "csv", the repository uses the `CSV_DATA_SOURCE_PATH` environment variable
to locate the CSV data source file.
If the `dbType` is "api", the repository utilizes the `API_DATA_SOURCE_KEY` and
`API_DATA_SOURCE_URL` environment variables to connect to the API.

Returns:
  - An implementation of the `DataStoreRepository` interface.
  - An error if the provided `dbType` value is unsupported.

Example:

	repo, err := InitDataStoreRepository("csv")
	if err != nil {
		log.Fatalf("Error initializing repository: %v", err)
	}
*/
func InitDataStoreRepository(dbType string) (DataStoreRepository, error) {
	switch dbType {
	case "csv":
		filePath := config.GetEnv("CSV_DATA_SOURCE_PATH")
		return NewCSVDatastoreRepository(filePath), nil

	case "api":
		apiKey := config.GetEnv("API_DATA_SOURCE_KEY")
		apiURL := config.GetEnv("API_DATA_SOURCE_URL")
		return NewApiDatastoreRepository(apiURL, apiKey), nil

	default:
		return nil, fmt.Errorf("unsupported DB_TYPE: %s", dbType)
	}
}
