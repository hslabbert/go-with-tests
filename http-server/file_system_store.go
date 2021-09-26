package main

import (
	"encoding/json"
	"io"
)

// A FileSystemPlayerStore implements the PlayerStore
// interface with filesystem backing, storing the
// Player data in a json file on disk.
type FileSystemPlayerStore struct {
	database io.Reader
}

// GetLeague returns a []Player of players stored in
// the provided *FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	_ = json.NewDecoder(f.database).Decode(&league)
	return league
}
