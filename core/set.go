package core

import "sync"

type SetDataType interface {
	comparable
}

type Set[T SetDataType] struct {
	sync.Mutex

	data map[T]bool
}

func NewSet[T SetDataType]() *Set[T] {
	return &Set[T]{
		data: make(map[T]bool),
	}
}

func (s *Set[T]) Add(item T) {
	s.Lock()
	defer s.Unlock()

	s.data[item] = true
}

func (s *Set[T]) Remove(item T) {
	s.Lock()
	defer s.Unlock()

	delete(s.data, item)
}

func (s *Set[T]) Contains(item T) bool {
	s.Lock()
	defer s.Unlock()

	_, ok := s.data[item]
	return ok
}
