package localstore

import "sync"

type MapStore[V any] struct {
	store map[string]V
	lock  sync.RWMutex
}

func newMapStore[T any]() MapStore[T] {
	return MapStore[T]{
		store: make(map[string]T),
	}
}

func (m *MapStore[V]) Set(key string, value V) {
	m.lock.Lock()
	m.store[key] = value
	m.lock.Unlock()
}

func (m *MapStore[V]) Get(key string) (V, bool) {
	m.lock.RLock()
	v, ok := m.store[key]
	m.lock.RUnlock()
	return v, ok
}

func (m *MapStore[V]) Delete(key string) {
	m.lock.Lock()
	delete(m.store, key)
	m.lock.Unlock()
}
