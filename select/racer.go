package racer

import (
	"net/http"
	"time"
)

// Racer takes two URLs as strings and returns the
// "faster" URL as a string. Racer will "race" the
// two URLs by sending an HTTP GET request to both of them
// and seeing which responds more quickly.
func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}
