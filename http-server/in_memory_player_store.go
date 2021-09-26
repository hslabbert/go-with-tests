package main

import "sync"

// NewInMemoryPlayerStore initializes and returns an
// *InMemoryPlayerStore in order to set up the needed
// map and mutex.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
		mu:    sync.RWMutex{}}
}

// An InMemoryPlayerStore is a PlayerStore that provides only
// ephemeral storage of player scores.
type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.RWMutex
}

// RecordWin adds one win to the provided player's score in the
// provided *InMemoryPlayerStore.
func (i *InMemoryPlayerStore) RecordWin(name string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
	return nil
}

// GetPlayerScore returns a player's score from the provided
// *InMemoryPlayerStore.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.store[name]
}

// GetLeague returns a League representing all of the players
// in the *InMemoryPlayerStore with their scores.
func (i *InMemoryPlayerStore) GetLeague() League {
	var league League
	i.mu.RLock()
	defer i.mu.RUnlock()
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
