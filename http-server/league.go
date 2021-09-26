package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// NewLeague takes a league of players provided as a json stream
// io.Reader and returns a []Player slice.
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
