package mocks

import (
	"sync"
	"time"
)

// MockRateLimiter for simulating rate-limiting behavior
type MockRateLimiter struct {
	AllowFunc func() bool
}

func (m *MockRateLimiter) Allow() bool {
	if m.AllowFunc != nil {
		return m.AllowFunc()
	}
	return true // Default behavior
}

type MockRateLimiter2 struct {
	Requests      int
	LimitPerSec   int
	LastResetTime time.Time
	Mu            sync.Mutex
	//AllowFunc func() bool
}

func (m *MockRateLimiter2) Allow() bool {
	//if m.AllowFunc != nil {
	//	return m.AllowFunc()
	//}
	return true // Default behavior
}

//
//// MockDataStore for simulating datastore interactions
//type MockDataStore struct {
//	forceErr bool
//}
//
//func (m *MockDataStore) Query(ip string) (models.IPInfo, error) {
//	if m.forceErr {
//		return models.IPInfo{}, errors.New("datastore error")
//	}
//	return models.IPInfo{Country: "United States", City: "New York"}, nil
//}
