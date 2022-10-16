package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70,
	}
	fmt.Println("Number of stocks: ", len(stocks))
	fmt.Println("Value of Microsoft: ", stocks["MSFT"])

	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	stocks["TSLA"] = 822.12
	fmt.Println("TSLA was added and it's value is: ", stocks["TSLA"])

	delete(stocks, "AMZN")
	fmt.Println("stocks minus AMZN: ", stocks)

	for key := range stocks {
		fmt.Println(key)
	}

	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}