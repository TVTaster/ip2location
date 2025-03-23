package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"ip2location/internal/models"
	"log"
	"net/http"
)

type apiDatastoreRepository struct {
	apiURL string
	apiKey string
}

func NewApiDatastoreRepository(apiURL string, apiKey string) DataStoreRepository {
	return &apiDatastoreRepository{apiKey: apiKey, apiURL: apiURL}
}

func (r *apiDatastoreRepository) GetLocationByIP(ip string) (*models.IPInfo, error) {
	// Construct the request URL
	url := fmt.Sprintf("%s?key=%s&ip=%s", r.apiURL, r.apiKey, ip)

	// Make the API request
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("failed to call IP2Location API: %w", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err)
		}
	}(resp.Body)

	// Decode the API response
	var apiResponse models.APIResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&apiResponse); err != nil {
		err = fmt.Errorf("failed to decode API response: %w", err)
		info := apiResponse.ToIPInfo()
		return &info, err
	}

	info := apiResponse.ToIPInfo()
	return &info, nil
}
