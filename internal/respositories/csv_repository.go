package repositories

import (
	"ip2location/internal/models"
	"ip2location/internal/services"
	"log"
)

type csvDatastoreRepository struct {
	store *models.MapDataStore
}

func NewCSVDatastoreRepository(filePath string) DataStoreRepository {
	store, err := services.NewCSVDataStore(filePath)
	if err != nil {
		log.Fatalf("failed to create CSV data store: %v", err)
		return nil
	}

	return &csvDatastoreRepository{store: store}
}

func (r *csvDatastoreRepository) GetLocationByIP(ip string) (*models.IPInfo, error) {
	find, err := r.store.GetLocationByIP(ip)
	if err != nil {
		log.Printf("an error occurred while finding IP %s: %v", ip, err)
		return nil, err
	}

	return find, nil
}
