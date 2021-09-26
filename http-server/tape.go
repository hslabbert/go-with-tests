package main

import (
	"fmt"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	err = t.file.Truncate(0)
	if err != nil {
		return 0, fmt.Errorf("problem truncating file %s, %v", t.file.Name(), err)
	}

	_, err = t.file.Seek(0, 0)
	if err != nil {
		return 0, fmt.Errorf("problem resetting to 0 offset on file %s, %v", t.file.Name(), err)
	}

	return t.file.Write(p)
}
