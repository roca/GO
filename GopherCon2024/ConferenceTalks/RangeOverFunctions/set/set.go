package set

import (
	"fmt"
	"iter"
)

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

	go func() {
		defer close(ch)
		for v := range s.m {
			select {
			case ch <- v:
			case <-stopCh:
				return
			}
		}
	}()

	next := func() (E, bool) {
		v, ok := <-ch
		return v, ok
	}

	stop := func() {
		close(stopCh)
	}

	return next, stop
}

func PrintAllElementsPull[E comparable](s *Set[E]) {
	next, stop := s.Pull()
	defer stop()
	for v, ok := next(); ok; v, ok = next() {
		fmt.Println(v)
	}
}

func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

func PrintAllElements[E comparable](s *Set[E]) {
	iterFunc := s.All()

	f := func(v E) bool {
		if s.Contains(v) {
			fmt.Println(v)
			return true
		}
		return false
	}

	// Call the function by passing the function as an argument
	// iterFunc(f)

	// Call the function by passing the function as a closure
	for v := range iterFunc {
		f(v)
	}
}

func EqSeq[E comparable](s1, s2 iter.Seq[E]) bool {
	next1, stop1 := iter.Pull(s1)
	defer stop1()
	next2, stop2 := iter.Pull(s2)
	defer stop2()

	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if !ok1 {
			return !ok2
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
}

func Filter[V any](f func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func PrintOddElements(s *Set[int]) {

	iterAll := s.All()

	f := func(v int) bool {
		if s.Contains(v) {
			fmt.Println(v)
			return true
		}
		return false
	}

	filterOddFunc := func(v int) bool {
		if v%2 != 0 {
			return true
		}
		return false
	}

	iterOdd := Filter(filterOddFunc, iterAll)

	for v := range iterOdd {
		f(v)
	}
}
