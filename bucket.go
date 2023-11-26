package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type BucketConfig struct {
	RefillRate  time.Duration
	RefillCount int // How much tokens should be periodically added to the bucket
	BucketSize  int
}

type BucketValue struct {
	LastRefillUtc   time.Time
	RemainingTokens int
}

type BucketLimiter struct{}

func NewBucketLimiter(storageProvider storage.Storage[BucketConfig, BucketValue]) *RateLimiter[BucketConfig, BucketValue] {
	return &RateLimiter[BucketConfig, BucketValue]{
		store:           storageProvider,
		InternalLimiter: &BucketLimiter{},
	}
}

func (b *BucketLimiter) TryAllow(count int, config BucketConfig, userValue BucketValue, nowUtc time.Time) (bool, BucketValue) {
	// Compute current bucket fill
	sinceLastRefill := nowUtc.Sub(userValue.LastRefillUtc)
	refillTimes := int(sinceLastRefill.Seconds() / config.RefillRate.Seconds())
	newFill := userValue.RemainingTokens + refillTimes*config.RefillCount
	if newFill > config.BucketSize {
		newFill = config.BucketSize
	}

	// Try allow count
	allow := newFill >= count
	if allow {
		newFill -= count
	}

	// Update user value
	userValue.RemainingTokens = newFill
	userValue.LastRefillUtc = userValue.LastRefillUtc.Add(time.Duration(refillTimes) * config.RefillRate) // Not nowUtc because it would make the refill time drift
	return allow, userValue
}
