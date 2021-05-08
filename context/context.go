package context

import (
	"fmt"
	"net/http"
)

// A Store can Fetch things.
type Store interface {
	Fetch() string
}

// Server takes a Store and returns an http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
