package dictionary

// A set of DictionaryErr definitions
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

// A DictionaryErr encapsulates various errors on accessing or
// manipulating dictionaries.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
