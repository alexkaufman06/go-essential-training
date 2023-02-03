package main

import (
	"fmt"
	"net/http"
	"time"
)

func returnType2(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s\n", url, err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
		"https://linkedin.com",
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go returnType2(url, ch)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}

	fmt.Println(time.Since(start))
}
