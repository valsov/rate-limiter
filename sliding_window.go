package rlimit

import (
	"time"

	"github.com/valsov/rlimit/data"
)

type SlidingWindowConfig struct {
	Start      time.Time
	WindowSize time.Duration
	MaxTokens  int
}

type SlidingWindowLimiter struct {
	BaseLimiter[SlidingWindowConfig, int] // int: Current tokens
}

func NewSlidingWindowLimiter(storageProvider data.Storage[SlidingWindowConfig, int], lockProvider data.Locker) Limiter {
	return &SlidingWindowLimiter{
		BaseLimiter: BaseLimiter[SlidingWindowConfig, int]{
			store:  storageProvider,
			locker: lockProvider,
		},
	}
}

func (s *SlidingWindowLimiter) TryAllow(id string, count int) (bool, error) {
	panic("TODO")
}
