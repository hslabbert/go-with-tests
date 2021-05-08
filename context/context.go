package context

import (
	"fmt"
	"net/http"
)

// A Store can Fetch things.
type Store interface {
	Fetch() string
	Cancel()
}

// Server takes a Store and returns an http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
