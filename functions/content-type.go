package main

import (
	"fmt"
	"http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR:  %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}


func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(resp)
	}
	return url, nil
}

