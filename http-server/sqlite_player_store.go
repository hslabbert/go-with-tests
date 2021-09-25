package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NewSqlitePlayerStore opens the specified sqlite3 database
// file and returns a *SqlitePlayerStore and any errors.
func NewSqlitePlayerStore(f string) (*SqlitePlayerStore, error) {
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		return nil, err
	}
	insertScoreStatement, err := db.Prepare("INSERT INTO playerscores(score, name) VALUES (?,?)")
	if err != nil {
		return nil, err
	}
	updateScoreStatement, err := db.Prepare("UPDATE playerscores SET score=? WHERE name=?")
	if err != nil {
		return nil, err
	}
	deleteScoreStatement, err := db.Prepare("DELETE from playerscores")
	if err != nil {
		return nil, err
	}

	return &SqlitePlayerStore{
		db:                   db,
		insertScoreStatement: insertScoreStatement,
		updateScoreStatement: updateScoreStatement,
		deleteScoreStatement: deleteScoreStatement,
	}, nil
}

// A SqlitePlayerStore is a PlayerStore with sqlite-backed storage
// It encapsulates a sqlite *sql.DB with a set of necessary prepared
// *sql.Stmt to perform various operations on the *sql.DB.
type SqlitePlayerStore struct {
	db                   *sql.DB
	insertScoreStatement *sql.Stmt
	updateScoreStatement *sql.Stmt
	deleteScoreStatement *sql.Stmt
}

// GetPlayerScore returns a player's score from the provided
// *SqlitePlayerStore.
func (s *SqlitePlayerStore) GetPlayerScore(name string) int {
	row := s.db.QueryRow("SELECT score FROM playerscores WHERE name=?", name)
	var score int
	err := row.Scan(&score)
	if err == sql.ErrNoRows {
		score = 0
	}
	return score
}

// RecordWin adds one win to the provided player's score in the
// provided *SqlitePlayerStore.
func (s *SqlitePlayerStore) RecordWin(name string) error {
	currentScore := s.GetPlayerScore(name)
	var statement *sql.Stmt
	if currentScore == 0 {
		statement = s.insertScoreStatement
	} else {
		statement = s.updateScoreStatement
	}
	newScore := currentScore + 1
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Stmt(statement).Exec(newScore, name)
	if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	return err
}

// DeletePlayerScores removes all player scores from the provided
// *SqlitePlayerStore in order to blank the score database.
func (s *SqlitePlayerStore) DeletePlayerScores() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Stmt(s.deleteScoreStatement).Exec()
	if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	return err
}

// GetLeague returns a []Player representing all of the players
// in the *SqlitePlayerStore with their scores.
func (s *SqlitePlayerStore) GetLeague() []Player {
	var league []Player
	rows, _ := s.db.Query("SELECT name, score FROM playerscores")
	var name string
	var wins int
	for rows.Next() {
		_ = rows.Scan(&name, &wins)
		league = append(league, Player{name, wins})
	}
	return league
}
