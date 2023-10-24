package data

import (
	"errors"
	"sync"
)

type Locker interface {
	Lock(id string) error
	Unlock(id string) error
}

type memoryLocker struct {
	locks sync.Map
}

func NewMemoryLocker() Locker {
	return &memoryLocker{sync.Map{}}
}

func (m *memoryLocker) Lock(id string) error {
	lock, _ := m.locks.LoadOrStore(id, &sync.Mutex{})
	lock.(*sync.Mutex).Lock()
	return nil
}

func (m *memoryLocker) Unlock(id string) error {
	lock, found := m.locks.Load(id)
	if !found {
		return errors.New("lock not found")
	}
	lock.(*sync.Mutex).Unlock()
	return nil
}
