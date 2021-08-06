package syncx

import "sync"

type Map struct {
	mu sync.RWMutex
	m  map[interface{}]interface{}
}

func NewMap() *Map {
	return &Map{
		mu: sync.RWMutex{},
		m:  map[interface{}]interface{}{},
	}
}

func (m *Map) Set(k, v interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[k] = v
}

func (m *Map) Get(k interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.m[k]
	return v, ok
}
