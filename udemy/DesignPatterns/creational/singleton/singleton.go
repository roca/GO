package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *singleton

func GetInstance() Singleton { // Return object must match interface
	if instance == nil {
		instance = new(singleton) // new() returns a pointer
	}
	return instance
}
