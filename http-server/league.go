package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

// A League is a []Player slice.
type League []Player

// NewLeague takes a league of players provided as a json stream
// io.Reader and returns a League slice.
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

// Find searches a League for the provided name and returns
// its Player record.
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
