package data

import "sync"

type Locker interface {
	Lock(id string)
	Unlock(id string)
}

type memoryLocker struct {
	locks sync.Map
}

func NewMemoryLocker() Locker {
	return &memoryLocker{sync.Map{}}
}

func (m *memoryLocker) Lock(id string) {
	lock, _ := m.locks.LoadOrStore(id, &sync.Mutex{})
	lock.(*sync.Mutex).Lock()
}

func (m *memoryLocker) Unlock(id string) {
	lock, found := m.locks.Load(id)
	if !found {
		return
	}
	lock.(*sync.Mutex).Unlock()
}
