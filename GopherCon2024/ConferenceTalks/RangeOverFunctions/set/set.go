package set

import "fmt"

// Set holds a set of elements

type Set[E comparable] struct {
	m map[E]struct{}
}

// New returns s new [Set]
func New[E comparable]() *Set[E] {
	return new(Set[E])
}

func (s *Set[E]) Add(v E) {
	if s.m == nil {
		s.m = make(map[E]struct{})
	}
	s.m[v] = struct{}{}
}

// Contains returns true if the set contains the element v
func (s *Set[E]) Contains(v E) bool {
	_, ok := s.m[v]
	return ok
}

// Union returns a new set with all the elements in s and t
func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
	r := New[E]()
	for k := range s1.m {
		r.Add(k)
	}
	for k := range s2.m {
		r.Add(k)
	}
	return r
}

func (s *Set[E]) Push(f func(E) bool) {
	for v := range s.m {
		if !f(v) {
			return
		}
	}
}

func PrintAllElemetsPush[E comparable](s *Set[E]) {
	s.Push(func(v E) bool {
		fmt.Println(v)
		return true
	})
}

// Pull returns a next function that returns each
// element of s with bool for whether the value
// is valid. The stop function should be called
// when finished calling the next function.
func (s *Set[E]) Pull() (func() (E, bool), func()) {
	ch := make(chan E)
	stopCh := make(chan bool)

	// go func () {
	// 	defer close(ch)
	// 	for v := range s.m {
	// 		select {
	// 		case ch <- v:
	// 		case <-stopCh:
	// 			return
	// 		}
	// 	}
	// }()

	next := func() (E, bool) {
		v, ok := <-ch
		return v, ok
	}

	stop := func() {
		close(stopCh)
	}

	return next, stop
}
