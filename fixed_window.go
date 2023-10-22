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

type FixedWindowLimiter struct {
	BaseLimiter[FixedWindowConfig, int] // int: Current tokens
}

func NewFixedWindowLimiter(storageProvider data.Storage[FixedWindowConfig, int], lockProvider data.Locker) Limiter {
	return &FixedWindowLimiter{
		BaseLimiter: BaseLimiter[FixedWindowConfig, int]{
			store:  storageProvider,
			locker: lockProvider,
		},
	}
}

func (f *FixedWindowLimiter) TryAllow(id string, count int) (bool, error) {
	panic("TODO")
}
