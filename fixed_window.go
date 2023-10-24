package rlimit

import (
	"time"

	"github.com/valsov/rlimit/data"
)

type FixedWindowConfig struct {
	Start      time.Time
	WindowSize time.Duration
	MaxTokens  int
}

type FixedWindowLimiter struct{}

func NewFixedWindowLimiter(storageProvider data.Storage[FixedWindowConfig, int], lockProvider data.Locker) *RateLimiter[FixedWindowConfig, int] {
	return &RateLimiter[FixedWindowConfig, int]{
		store:   storageProvider,
		locker:  lockProvider,
		Limiter: &FixedWindowLimiter{},
	}
}

func (f *FixedWindowLimiter) ValidateConfig(config FixedWindowConfig) error {
	panic("TODO")
}

func (f *FixedWindowLimiter) TryAllow(count int, config FixedWindowConfig, userValue int) bool {
	panic("todo")
}
