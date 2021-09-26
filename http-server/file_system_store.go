package main

import (
	"encoding/json"
	"os"
)

// A FileSystemPlayerStore implements the PlayerStore
// interface with filesystem backing, storing the
// Player data in a json file on disk.
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore constructs a *FileSystemPlayerStore with the
// provided json-formatted database file.
func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}

// GetLeague returns a League of players stored in
// the provided *FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

// GetPlayerScore retrieves the current score for the provided
// player from the provided *FileSystemPlayerStore.
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin increments the score of the named player in the
// provided *FileSystemPlayerStore.
func (f *FileSystemPlayerStore) RecordWin(name string) error {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	err := f.database.Encode(f.league)
	return err
}
