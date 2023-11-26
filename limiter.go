package rlimit

import (
	"time"

	"github.com/valsov/rlimit/storage"
)

type Limiter[TConfig, TValue any] interface {
	TryAllow(count int, config TConfig, userValue TValue, now time.Time) bool
}

type RateLimiter[TConfig, TValue any] struct {
	store           storage.Storage[TConfig, TValue]
	GlobalConfig    TConfig
	InternalLimiter Limiter[TConfig, TValue]
}

func (r *RateLimiter[TConfig, TValue]) GlobalConfigure() error {
	panic("TODO")
}

func (r *RateLimiter[TConfig, TValue]) Configure(id string) error {
	panic("TODO")
}

func (r *RateLimiter[TConfig, TValue]) TryAllow(id string, count int) (allowed bool, err error) {
	data, err := r.store.Get(id)
	if err != nil {
		return false, nil
	}
	config := data.Config
	if !data.HasConfig {
		config = r.GlobalConfig
	}

	allowed = r.InternalLimiter.TryAllow(count, config, data.Value, time.Now())

	err = r.store.Set(id, data)
	return allowed, err
}
