package rlimit

type Storage[T any] interface {
	Get(id string) T
	Set(id string, data T) error
}

type memoryStore[T any] struct {
	store map[string]T
}

func NewMemoryStore[T any]() Storage[T] {
	return &memoryStore[T]{make(map[string]T)}
}

func (i *memoryStore[T]) Get(id string) T {
	panic("TODO")
}

func (i *memoryStore[T]) Set(id string, data T) error {
	panic("TODO")
}
