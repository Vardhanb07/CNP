package base

import (
	"sync"
)

type Store[T any] struct {
	items []T
	limit int
	mu    sync.Mutex
}

func NewStore[T any](limit int) *Store[T] {
	return &Store[T]{
		limit: limit,
	}
}

func (s *Store[T]) Set(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) > s.limit {
		return false
	}
	s.items = append(s.items, item)
	return true
}

func (s *Store[T]) Get() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.items
}

func (s *Store[T]) At(idx int) (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var def T
	if 0 > idx && idx >= len(s.items) {
		return def, false
	}
	return s.items[idx], true
}

func (s *Store[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}

func (s *Store[T]) Discard() {
	s.mu.Lock()
	defer s.mu.Unlock()
	var def []T
	s.items = def
}
