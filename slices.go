package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)
	fmt.Println("Length: ", len(loons))
	fmt.Println(loons[1:])
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	loons = append(loons, "elmer")
	for _, name := range loons {
		fmt.Println(name)
	}
}
