/*
An "even ended number" is a number whose first and last digits are the same

Count how many "even ended numbers" are there that are multiplication of two 4 digit numbers.
 */

package main

import (
	"fmt"
)

func main() {
	count := 0

	for a:= 1000; a <= 9999; a++ {
		for b := a; b <=9999; b++ {
			n := a * b

			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s) - 1] {
				count++
			}
			//fmt.Println("Is even ended number: ", s, s[0] == s[len(s) - 1])
		}
	}

	fmt.Println("The total number of even ended numbers is: ", count)


	//n := 404
	//s := fmt.Sprintf("%d", n)
	//
	//fmt.Println("s = %q (type %T)\n", s, s)
	//fmt.Println("Is even ended number: ", s[0] == s[len(s) - 1])
	//fmt.Println(s[len(s) - 1])
}