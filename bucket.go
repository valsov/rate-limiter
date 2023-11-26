package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type BucketConfig struct {
	RefillRate time.Duration
	BucketSize int
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

func (b *BucketLimiter) TryAllow(count int, config BucketConfig, userValue BucketValue, now time.Time) bool {
	panic("todo")
}
