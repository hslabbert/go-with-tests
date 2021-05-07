package concurrency

// A WebsiteChecker takes a URL as a string and returns a bool
// for whether it passed its check.
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites takes a WebsiteChecker and a slice of URLs to check,
// runs them through the WebsiteChecker, and returns a map of
// the URLs and whether they passed checks (as bools).
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
