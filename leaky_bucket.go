package rlimit

import (
	"time"

	"github.com/valsov/rlimit/data"
)

type LeakyBucketConfig struct {
	FlowTick      time.Duration
	FlowTickCount int
	BucketSize    int
}

type LeakyBucketLimiter struct{}

func NewLeakyBucketLimiter(storageProvider data.Storage[LeakyBucketConfig, int], lockProvider data.Locker) *RateLimiter[LeakyBucketConfig, int] {
	return &RateLimiter[LeakyBucketConfig, int]{
		store:   storageProvider,
		locker:  lockProvider,
		Limiter: &LeakyBucketLimiter{},
	}
}

func (l *LeakyBucketLimiter) ValidateConfig(config LeakyBucketConfig) error {
	panic("TODO")
}

func (l *LeakyBucketLimiter) TryAllow(count int, config LeakyBucketConfig, userValue int) bool {
	panic("todo")
}
