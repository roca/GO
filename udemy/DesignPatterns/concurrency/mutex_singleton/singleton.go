package mutex_singleton

import "sync"

type ISingleton interface {
	AddOne()
	GetCount() int
}

type singleton struct {
	count int
	sync.RWMutex
}

func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}

var instance *singleton

func GetInstance() ISingleton { // Return object must match interface
	if instance == nil {
		instance = new(singleton) // new() returns a pointer
	}
	return instance
}
