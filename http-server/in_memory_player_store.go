package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
		mu:    sync.RWMutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.RWMutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.store[name]
}
