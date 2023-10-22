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

type LeakyBucketLimiter struct {
	BaseLimiter[LeakyBucketConfig, int] // int: Current filling
}

func NewLeakyBucketLimiter(storageProvider data.Storage[LeakyBucketConfig, int], lockProvider data.Locker) Limiter {
	return &LeakyBucketLimiter{
		BaseLimiter: BaseLimiter[LeakyBucketConfig, int]{
			store:  storageProvider,
			locker: lockProvider,
		},
	}
}

func (l *LeakyBucketLimiter) TryAllow(id string, count int) (bool, error) {
	panic("TODO")
}
