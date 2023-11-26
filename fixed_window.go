package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type FixedWindowConfig struct {
	WindowLength time.Duration
	MaxTokens    int
}

type FixedWindowValue struct {
	WindowStartUtc  time.Time
	RemainingTokens int
}

type FixedWindowLimiter struct{}

func NewFixedWindowLimiter(storageProvider storage.Storage[FixedWindowConfig, FixedWindowValue]) *RateLimiter[FixedWindowConfig, FixedWindowValue] {
	return &RateLimiter[FixedWindowConfig, FixedWindowValue]{
		store:           storageProvider,
		InternalLimiter: &FixedWindowLimiter{},
	}
}

func (f *FixedWindowLimiter) TryAllow(count int, config FixedWindowConfig, userValue FixedWindowValue, nowUtc time.Time) (bool, FixedWindowValue) {
	sinceLastWindow := nowUtc.Sub(userValue.WindowStartUtc)
	if sinceLastWindow >= config.WindowLength {
		userValue.WindowStartUtc = nowUtc
		userValue.RemainingTokens = config.MaxTokens
	}

	if userValue.RemainingTokens < count {
		return false, userValue
	}

	userValue.RemainingTokens -= count
	return true, userValue
}
