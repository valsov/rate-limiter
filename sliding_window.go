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

type SlidingWindowLimiter struct{}

func NewSlidingWindowLimiter(storageProvider data.Storage[SlidingWindowConfig, int], lockProvider data.Locker) *RateLimiter[SlidingWindowConfig, int] {
	return &RateLimiter[SlidingWindowConfig, int]{
		store:   storageProvider,
		locker:  lockProvider,
		Limiter: &SlidingWindowLimiter{},
	}
}

func (s *SlidingWindowLimiter) ValidateConfig(config SlidingWindowConfig) error {
	panic("TODO")
}

func (s *SlidingWindowLimiter) TryAllow(count int, config SlidingWindowConfig, userValue int) bool {
	panic("todo")
}
