package models

import "fmt"

type IPInfo struct {
	City    string
	Country string
}

type APIResponse struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionName  string  `json:"region_name"`
	CityName    string  `json:"city_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	ASN         string  `json:"asn"`
	AS          string  `json:"as"`
	IsProxy     bool    `json:"is_proxy"`
}

func (r APIResponse) ToIPInfo() IPInfo {
	return IPInfo{
		City:    r.CityName,
		Country: r.CountryName,
	}
}

type DataStore interface {
	GetLocationByIP(ip string) (*IPInfo, error)
}

type MapDataStore struct {
	Data map[string]IPInfo
}

func (ds *MapDataStore) GetLocationByIP(ip string) (*IPInfo, error) {
	info, ok := ds.Data[ip]
	if !ok {
		return &IPInfo{}, fmt.Errorf("IP not found")
	}
	return &info, nil
}
