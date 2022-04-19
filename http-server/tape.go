package poker

import (
	"fmt"
	"os"
)

// A Tape reads a file on disk and provides seek functionality.
type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	err = t.File.Truncate(0)
	if err != nil {
		return 0, fmt.Errorf("problem truncating file %s, %v", t.File.Name(), err)
	}

	_, err = t.File.Seek(0, 0)
	if err != nil {
		return 0, fmt.Errorf("problem resetting to 0 offset on file %s, %v", t.File.Name(), err)
	}

	return t.File.Write(p)
}
