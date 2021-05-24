package main

import (
	"os"
	"time"

	"github.com/hslabbert/go-with-tests/math/pkg/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
