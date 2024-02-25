package rlimit

import (
	"errors"
	"time"

	"github.com/valsov/rlimit/storage"
)

type Limiter[TConfig, TValue any] interface {
	TryAllow(count int, config TConfig, userValue TValue, nowUtc time.Time) (bool, TValue)
}

type RateLimiter[TConfig, TValue any] struct {
	store           storage.Storage[TConfig, TValue]
	GlobalConfig    TConfig
	InternalLimiter Limiter[TConfig, TValue]
}

func (r *RateLimiter[TConfig, TValue]) GlobalConfigure(config TConfig) {
	r.GlobalConfig = config
}

func (r *RateLimiter[TConfig, TValue]) Configure(id string, config TConfig) error {
	data, err := r.store.Get(id)
	if err != nil {
		return err
	}
	data.Config = config
	return r.store.Set(id, data)
}

func (r *RateLimiter[TConfig, TValue]) TryAllow(id string, count int) (bool, error) {
	if count <= 0 {
		return false, errors.New("count must be greater than 0")
	}

	data, err := r.store.Get(id)
	if err != nil {
		return false, nil
	}
	config := data.Config
	if !data.HasConfig {
		config = r.GlobalConfig
	}

	allowed, newUserValue := r.InternalLimiter.TryAllow(count, config, data.Value, time.Now().UTC())
	data.Value = newUserValue

	err = r.store.Set(id, data)
	return allowed, err
}
