package rlimit

import "sync"

type Storage[T any] interface {
	Get(id string) T
	Set(id string, data T) error
}

type Locker interface {
	RLock(id string)
	RUnlock(id string)
	WLock(id string)
	WUnlock(id string)
}

type memoryStore[T any] struct {
	store map[string]T
}

func NewMemoryStore[T any]() Storage[T] {
	return &memoryStore[T]{make(map[string]T)}
}

func (s *memoryStore[T]) Get(id string) T {
	return s.store[id]
}

func (s *memoryStore[T]) Set(id string, data T) error {
	s.store[id] = data
	return nil
}

type memoryLocker struct {
	locks sync.Map
}

func NewMemoryLocker() Locker {
	return &memoryLocker{sync.Map{}}
}

func (m *memoryLocker) RLock(id string) {
	lock, _ := m.locks.LoadOrStore(id, &sync.RWMutex{})
	lock.(*sync.RWMutex).RLock()
}

func (m *memoryLocker) RUnlock(id string) {
	lock, found := m.locks.Load(id)
	if !found {
		return
	}
	lock.(*sync.RWMutex).RUnlock()
}

func (m *memoryLocker) WLock(id string) {
	lock, _ := m.locks.LoadOrStore(id, &sync.RWMutex{})
	lock.(*sync.RWMutex).Lock()
}

func (m *memoryLocker) WUnlock(id string) {
	lock, found := m.locks.Load(id)
	if !found {
		return
	}
	lock.(*sync.RWMutex).Unlock()
}
