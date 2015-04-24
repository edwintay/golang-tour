package main

import (
  "fmt"
  "sync"
  "log"
  "os"
)

type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}

var visited = struct {
  urls map[string]bool
  mutex sync.Mutex
}{urls: make(map[string]bool)}

var logger = log.New(os.Stdout, "", 0)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  crawl(url, depth, fetcher)
}

// helper to Crawl that returns error status
func crawl(url string, depth int, fetcher Fetcher) error {
  if depth <= 0 {
    return fmt.Errorf("max depth reached")
  }
  visited.mutex.Lock()
  if _, ok := visited.urls[url]; ok {
    visited.mutex.Unlock()
    return fmt.Errorf("already fetched: %s", url)
  }
  visited.urls[url] = true // record visit
  visited.mutex.Unlock()

  body, urls, err := fetcher.Fetch(url)
  if err != nil {
    return err
  }
  logger.Printf("found: %s %q\n", url, body)
  done := make(chan string, len(urls))
  for i, u := range urls {
    logger.Printf("-> [%s] crawling child %d/%d: %s\n",
                url, i, len(urls), u)
    go func(u string) {
      err := crawl(u, depth-1, fetcher)
      if err != nil {
        logger.Printf("[%s] %v\n", url, err)
      }
      done <- u
    }(u)
  }
  for i := 0; i < len(urls); i++ {
    logger.Printf("<- [%s] waiting for child %d/%d\n", url, i, len(urls))
    u := <- done
    logger.Printf("<- done with %s\n", u)
  }
  return nil
}

func main() {
  Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
  "http://golang.org/": &fakeResult{
    "The Go Programming Language",
    []string{
      "http://golang.org/pkg/",
      "http://golang.org/cmd/",
    },
  },
  "http://golang.org/pkg/": &fakeResult{
    "Packages",
    []string{
      "http://golang.org/",
      "http://golang.org/cmd/",
      "http://golang.org/pkg/fmt/",
      "http://golang.org/pkg/os/",
    },
  },
  "http://golang.org/pkg/fmt/": &fakeResult{
    "Package fmt",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
  "http://golang.org/pkg/os/": &fakeResult{
    "Package os",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
}

