// count how many times a word appears in text
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "This is the greatest and best program in the world and this is awesome"
	words := strings.Fields(text)
	counts := map[string]int{}
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}