package main

import (
	"io"
)

// A FileSystemPlayerStore implements the PlayerStore
// interface with filesystem backing, storing the
// Player data in a json file on disk.
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague returns a []Player of players stored in
// the provided *FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
