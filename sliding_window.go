package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type SlidingWindowConfig struct {
	Start      time.Time
	WindowSize time.Duration
	MaxTokens  int
}

type SlidingWindowLimiter struct{}

func NewSlidingWindowLimiter(storageProvider storage.Storage[SlidingWindowConfig, int]) *RateLimiter[SlidingWindowConfig, int] {
	return &RateLimiter[SlidingWindowConfig, int]{
		store:           storageProvider,
		InternalLimiter: &SlidingWindowLimiter{},
	}
}

func (s *SlidingWindowLimiter) TryAllow(count int, config SlidingWindowConfig, userValue int, now time.Time) bool {
	panic("todo")
}
