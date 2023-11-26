package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type LeakyBucketConfig struct {
	NewTokensRate  time.Duration
	NewTokensCount int
	BucketSize     int
}

type LeakyBucketValue struct {
	LastCheckedUtc  time.Time
	RemainingTokens int
}

type LeakyBucketLimiter struct{}

func NewLeakyBucketLimiter(storageProvider storage.Storage[LeakyBucketConfig, LeakyBucketValue]) *RateLimiter[LeakyBucketConfig, LeakyBucketValue] {
	return &RateLimiter[LeakyBucketConfig, LeakyBucketValue]{
		store:           storageProvider,
		InternalLimiter: &LeakyBucketLimiter{},
	}
}

func (l *LeakyBucketLimiter) TryAllow(count int, config LeakyBucketConfig, userValue LeakyBucketValue, nowUtc time.Time) (bool, LeakyBucketValue) {
	// Compute current bucket fill
	sinceLastCheck := nowUtc.Sub(userValue.LastCheckedUtc)
	refillTimes := int(sinceLastCheck.Seconds() / config.NewTokensRate.Seconds())
	userValue.RemainingTokens += refillTimes * config.NewTokensCount
	if userValue.RemainingTokens > config.BucketSize {
		userValue.RemainingTokens = config.BucketSize
	}

	// Try allow count
	allow := userValue.RemainingTokens >= count
	if allow {
		userValue.RemainingTokens -= count
	}

	userValue.LastCheckedUtc = nowUtc
	return allow, userValue
}
