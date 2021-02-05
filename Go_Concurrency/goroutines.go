package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	var wg sync.WaitGroup // [ wait group ]
	urls := []string {
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	for _, url := range urls {
		wg.Add(1) // [ wait group ]
		go func(url string) { //run not wait for process
			returnType(url)
			wg.Done() // [ wait group ]
		} (url)
		// you can use sleep to wait for them (bad ways) 
		//or wait group, but we do have channels
	}
	wg.Wait() // [ wait group ]
}