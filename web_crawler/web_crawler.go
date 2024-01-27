package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	mu sync.Mutex
	v  map[string]bool
}

func (s *SafeMap) SetValue(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.v[key] = true
}

func (s *SafeMap) HasVisited(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.v[key]
	return ok
}

func Crawl(url string, depth int, fetcher Fetcher, visitedURLs *SafeMap) {
	if depth <= 0 {
		return
	}

	if visitedURLs.HasVisited(url) {
		return
	}

	visitedURLs.SetValue(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	ch := make(chan bool)

	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher, visitedURLs)
			ch <- true
		}(u)
	}

	for range urls {
		<-ch
	}
}

func main() {
	visitedURLs := &SafeMap{v: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, visitedURLs)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
