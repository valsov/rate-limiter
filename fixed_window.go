package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type FixedWindowConfig struct {
	Start      time.Time
	WindowSize time.Duration
	MaxTokens  int
}

type FixedWindowLimiter struct{}

func NewFixedWindowLimiter(storageProvider storage.Storage[FixedWindowConfig, int]) *RateLimiter[FixedWindowConfig, int] {
	return &RateLimiter[FixedWindowConfig, int]{
		store:           storageProvider,
		InternalLimiter: &FixedWindowLimiter{},
	}
}

func (f *FixedWindowLimiter) TryAllow(count int, config FixedWindowConfig, userValue int, nowUtc time.Time) (bool, int) {
	panic("todo")
}
