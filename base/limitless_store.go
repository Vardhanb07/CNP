package base

import "sync"

type LStore[T any] struct {
	items []T
	mu    sync.Mutex
}

func NewLStore[T any]() *LStore[T] {
	return &LStore[T]{}
}

func (ls *LStore[T]) Set(item T) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.items = append(ls.items, item)
}

func (ls *LStore[T]) Get() []T {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	return ls.items
}

func (ls *LStore[T]) At(idx int) (T, bool) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	var def T
	if 0 > idx || idx >= len(ls.items) {
		return def, false
	}
	return ls.items[idx], true
}

func (ls *LStore[T]) Len() int {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	return len(ls.items)
}

func (ls *LStore[T]) Discard() {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.items = []T{}
}
