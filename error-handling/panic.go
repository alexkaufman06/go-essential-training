package main

import "fmt"

func main() {
	values := []int{1, 2, 3}
	v, err := safeValue(values, 10)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	fmt.Println("v: ", v)

	//v := values[10] // This will cause a panic
	//fmt.Println(v)

	// panic("oops") // NOT COMMON TO DO THIS
}

func safeValue(vals []int, index int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return vals[index], nil
}
