package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type SlidingWindowConfig struct {
	WindowLength time.Duration
	MaxTokens    int
}

type SlidingWindowValue struct {
	PreviousConsumption int
	CurrentConsumption  int
	WindowStartUtc      time.Time
}

type SlidingWindowLimiter struct{}

func NewSlidingWindowLimiter(storageProvider storage.Storage[SlidingWindowConfig, SlidingWindowValue]) *RateLimiter[SlidingWindowConfig, SlidingWindowValue] {
	return &RateLimiter[SlidingWindowConfig, SlidingWindowValue]{
		store:           storageProvider,
		InternalLimiter: &SlidingWindowLimiter{},
	}
}

func (s *SlidingWindowLimiter) TryAllow(count int, config SlidingWindowConfig, userValue SlidingWindowValue, nowUtc time.Time) (bool, SlidingWindowValue) {
	sinceLastWindow := nowUtc.Sub(userValue.WindowStartUtc)
	if sinceLastWindow >= config.WindowLength {
		if sinceLastWindow >= config.WindowLength*2 {
			// Previous window is out of bounds
			userValue.PreviousConsumption = 0
		} else {
			userValue.PreviousConsumption = userValue.CurrentConsumption
		}
		userValue.CurrentConsumption = 0
		userValue.WindowStartUtc = nowUtc
	}

	previousWindowRatio := 1.0 - (nowUtc.Sub(userValue.WindowStartUtc).Seconds() / config.WindowLength.Seconds())
	consumption := int(previousWindowRatio*float64(userValue.PreviousConsumption)) + userValue.CurrentConsumption
	allow := consumption+count <= config.MaxTokens
	if allow {
		userValue.CurrentConsumption += count
	}

	return allow, userValue
}
