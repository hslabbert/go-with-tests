package poker_test

import (
	"io/ioutil"
	"testing"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &poker.Tape{file}

	_, err := tape.Write([]byte("abc"))
	assertNoError(t, err)

	_, err = file.Seek(0, 0)
	assertNoError(t, err)

	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
