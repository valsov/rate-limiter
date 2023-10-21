package rlimit

type Limiter interface {
	TryAllow(id string, count int) (bool, error)
}

type sampleLimiter struct {
	store  Storage[string]
	locker Locker
}

func NewSampleLimiter(store Storage[string], locker Locker) Limiter {
	return &sampleLimiter{store, locker}
}

func (s *sampleLimiter) TryAllow(id string, count int) (bool, error) {
	panic("TODO")
}
