package rlimit

import "github.com/valsov/rlimit/data"

type Limiter interface {
	GlobalInit() error
	Init(id string) error
	TryAllow(id string, count int) (bool, error)
}

type BaseLimiter[TConfig, TValue any] struct {
	store        data.Storage[TConfig, TValue]
	locker       data.Locker
	GlobalConfig TConfig
}

func (l *BaseLimiter[TConfig, TValue]) Init(id string) error {
	l.locker.Lock(id)
	defer l.locker.Unlock(id)

	panic("TODO")
}

func (l *BaseLimiter[TConfig, TValue]) GlobalInit() error {
	panic("TODO")
}
