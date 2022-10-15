// Calculate maximal value in slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]
	for _, num := range nums[1:] {
		if max < num {
			max = num
		}
	}
	fmt.Println("The maximal value in this slice is: ", max)
}