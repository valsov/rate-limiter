package rlimit

type Limiter interface {
	Try(id string, count int) (bool, error)
}

type sampleLimiter struct {
	store Storage[string]
}

func NewSampleLimiter(store Storage[string]) Limiter {
	return &sampleLimiter{store}
}

func (s *sampleLimiter) Try(id string, count int) (bool, error) {
	panic("TODO")
}
