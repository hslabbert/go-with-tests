package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing player DB file %s, %v", file.Name(), err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

// GetLeague returns a League of players stored in
// the provided *FileSystemPlayerStore
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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

func initializePlayerDBFile(file *os.File) error {
	_, err := file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("problem resetting to 0 offset on file %s, %v", file.Name(), err)
	}

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		_, err := file.Write([]byte("[]"))
		if err != nil {
			return fmt.Errorf("problem writing empty league to blank file %s, %v", file.Name(), err)
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			return fmt.Errorf("problem resetting to 0 offset on file %s, %v", file.Name(), err)
		}
	}

	return nil
}

// FileSystemPlayerStoreFromFile constructs a FileSystemPlayerStore from
// the provided file path, with the needed func() to close the
// FileSystemPlayerStore json db file.
func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store %v", err)
	}

	return store, closeFunc, err
}
