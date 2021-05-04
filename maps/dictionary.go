package dictionary

import "errors"

// ErrNotFound is returned if a definition is not found in a
// Dictionary for a word or phrase.
var ErrNotFound = errors.New("could not find the word you were looking for")

// A Dictionary is a map
type Dictionary map[string]string

// Search takes a word to look up in the dictionary
// and returns that word's definition.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add will add a word into an existing Dictionary.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
