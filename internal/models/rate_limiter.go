package models

import (
	"sync"
	"time"
)

type RateLimiterInterface interface {
	Allow() bool
}

type RateLimiter struct {
	Requests      int
	LimitPerSec   int
	LastResetTime time.Time
	Mu            sync.Mutex
}

func (r *RateLimiter) Allow() bool {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	now := time.Now()
	if now.Sub(r.LastResetTime) >= time.Second {
		r.Requests = 0
		r.LastResetTime = now
	}

	if r.Requests < r.LimitPerSec {
		r.Requests++
		return true
	}
	return false
}
