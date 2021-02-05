package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		out <- fmt.Sprintf("%s -> %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s\n", url, ctype)
}

func main() {

	urls := []string {
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}
	
	//var wg sync.WaitGroup // [ wait group ]

	/*
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
	*/
}