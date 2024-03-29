package poker_test

import (
	"io/ioutil"
	"os"
	"testing"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

func TestFileSystemStore(t *testing.T) {

	database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)

	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)

	assertNoError(t, err)

	t.Run("league from a reader", func(t *testing.T) {
		got := store.GetLeague()

		want := poker.League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		poker.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("league sorted", func(t *testing.T) {
		got := store.GetLeague()

		want := []poker.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		poker.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		err := store.RecordWin("Chris")
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		err := store.RecordWin("Pepper")
		assertNoError(t, err)

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	// NOTE: this test reinitializes an empty file rather than reusing the
	// same/existing database file from the earlier runners.
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase = createTempFile(t, "")
		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, err = tmpfile.Write([]byte(initialData))

	assertNoError(t, err)

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
