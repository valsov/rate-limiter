package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type LeakyBucketConfig struct {
	FlowTick      time.Duration
	FlowTickCount int
	BucketSize    int
}

type LeakyBucketLimiter struct{}

func NewLeakyBucketLimiter(storageProvider storage.Storage[LeakyBucketConfig, int]) *RateLimiter[LeakyBucketConfig, int] {
	return &RateLimiter[LeakyBucketConfig, int]{
		store:           storageProvider,
		InternalLimiter: &LeakyBucketLimiter{},
	}
}

func (l *LeakyBucketLimiter) TryAllow(count int, config LeakyBucketConfig, userValue int, nowUtc time.Time) (bool, int) {
	panic("todo")
}
