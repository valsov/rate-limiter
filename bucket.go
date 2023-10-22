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

type BucketLimiter struct {
	BaseLimiter[BucketConfig, BucketValue]
}

func NewBucketLimiter(storageProvider data.Storage[BucketConfig, BucketValue], lockProvider data.Locker) Limiter {
	return &BucketLimiter{
		BaseLimiter: BaseLimiter[BucketConfig, BucketValue]{
			store:  storageProvider,
			locker: lockProvider,
		},
	}
}

func (b *BucketLimiter) TryAllow(id string, count int) (bool, error) {
	panic("todo")
}
