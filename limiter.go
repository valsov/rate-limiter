package rlimit

import "github.com/valsov/rlimit/data"

type RateLimiter[TConfig, TValue any] struct {
	store        data.Storage[TConfig, TValue]
	locker       data.Locker
	GlobalConfig TConfig
	Limiter      Limiter[TConfig, TValue]
}

func (r *RateLimiter[TConfig, TValue]) GlobalConfigure() error {
	panic("TODO")
}

func (r *RateLimiter[TConfig, TValue]) Configure(id string) error {
	r.locker.Lock(id)
	defer r.locker.Unlock(id)

	panic("TODO")
}

func (r *RateLimiter[TConfig, TValue]) TryAllow(id string, count int) (allowed bool, err error) {
	err = r.locker.Lock(id)
	if err != nil {
		return false, nil
	}
	defer func() {
		unlockErr := r.locker.Unlock(id)
		if err == nil {
			err = unlockErr // Will be returned along with allowed boolean
		}
	}()

	data, err := r.store.Get(id)
	if err != nil {
		return false, nil
	}
	config := data.Config
	if !data.HasConfig {
		config = r.GlobalConfig
	}

	err = r.Limiter.ValidateConfig(config)
	if err != nil {
		return false, nil
	}

	allowed = r.Limiter.TryAllow(count, config, data.Value)

	err = r.store.Set(id, data)
	return allowed, err
}

type Limiter[TConfig, TValue any] interface {
	ValidateConfig(config TConfig) error
	TryAllow(count int, config TConfig, userValue TValue) bool
}
