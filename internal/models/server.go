package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	Port        string
	Datastore   DataStore
	RateLimiter *RateLimiter
}

func (s *Server) Start() {
	http.HandleFunc("/v1/find-country", s.findCountryByIpHandler)
	http.HandleFunc("/v1/health", s.healthHandler)

	log.Printf("Starting server on port %s", s.Port)
	err := http.ListenAndServe(":"+s.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if !s.RateLimiter.Allow() {
		log.Println("Rate limit exceeded")
		http.Error(w, `{"error":"Rate limit exceeded"}`, http.StatusTooManyRequests)
		return
	}

	_, err := w.Write([]byte("Hello, this is the IP2Location server!"))
	if err != nil {
		return
	}
}

func (s *Server) findCountryByIpHandler(w http.ResponseWriter, r *http.Request) {
	if !s.RateLimiter.Allow() {
		log.Println("Rate limit exceeded")
		http.Error(w, `{"error":"Rate limit exceeded"}`, http.StatusTooManyRequests)
		return
	}

	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, `{"error":"Missing IP parameter"}`, http.StatusBadRequest)
		return
	}

	info, err := s.Datastore.GetLocationByIP(ip)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
		return
	}

	response := map[string]string{
		"country": info.Country,
		"city":    info.City,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Failed to encode response: %v", err)
	}
}
