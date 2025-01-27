package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://hotstar.com",
		"https://jio.com",
		"https://flipkart.com",
	}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		url := url
		go func() {
			defer wg.Done()
			siteTime(url)
		}()
	}
	wg.Wait()
}

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Print("error happened", err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Print("some error")
	}

	log.Printf("Time took for %s is %d", url , time.Since(start))
}
