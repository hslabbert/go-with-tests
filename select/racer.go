package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer is takes two URLs as strings and returns the
// "faster" URL as a string, with a 10 second timeout.
// Racer will "race" the two URLs by sending an HTTP GET
// request to both of them and seeing which responds more quickly.
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer is takes two URLs as strings and a
// time.Duration timeout, and returns the "faster" URL as a string.
// ConfigurableRacer will "race" the two URLs by sending an HTTP GET
// request to both of them and seeing which responds more quickly.
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
