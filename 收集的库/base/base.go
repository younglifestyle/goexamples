package base

import (
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	data   map[string]bool
}

func (s *SafeMap) Set(verstr string) {
	s.Lock()
	defer s.Unlock()
	s.data[verstr] = true
}

func (s *SafeMap) Get(verstr string) bool {
	s.RLock()
	defer s.RUnlock()
	return s.data[verstr]
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data:   make(map[string]bool),
	}
}
