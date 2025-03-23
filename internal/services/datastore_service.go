package services

import (
	"encoding/csv"
	"ip2location/internal/models"
	"log"
	"os"
)

func NewCSVDataStore(filePath string) (*models.MapDataStore, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	data := make(map[string]models.IPInfo)
	for _, record := range records {
		if len(record) != 3 {
			continue
		}
		ip := record[0]
		city := record[1]
		country := record[2]
		data[ip] = models.IPInfo{City: city, Country: country}
	}
	log.Println("Loaded CSV data successfully")
	return &models.MapDataStore{Data: data}, nil
}
