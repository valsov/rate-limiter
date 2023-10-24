package rlimit

import (
	"time"

	"github.com/valsov/rlimit/data"
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

func NewBucketLimiter(storageProvider data.Storage[BucketConfig, BucketValue], lockProvider data.Locker) *RateLimiter[BucketConfig, BucketValue] {
	return &RateLimiter[BucketConfig, BucketValue]{
		store:   storageProvider,
		locker:  lockProvider,
		Limiter: &BucketLimiter{},
	}
}

func (b *BucketLimiter) ValidateConfig(config BucketConfig) error {
	panic("TODO")
}

func (b *BucketLimiter) TryAllow(count int, config BucketConfig, userValue BucketValue) bool {
	panic("todo")
}
