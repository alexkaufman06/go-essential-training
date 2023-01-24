package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignId string
	Balance    float64 // USD
	Expires    time.Time
}

func main() {
	b1 := Budget{
		CampaignId: "Give rich people more money",
		Balance:    1000000,
		Expires:    time.Now().Add(7 * 24 * time.Hour),
	}

	fmt.Println("Print the contents of the struct:")
	fmt.Println(b1)
	fmt.Println("Print contents of scruct with property information:")
	fmt.Printf("%#v\n", b1)
	fmt.Println("Print the CampaignId:")
	fmt.Println(b1.CampaignId)

	b2 := Budget{
		CampaignId: "Save the poor",
		Balance:    25,
	}

	fmt.Println("Print b2 value with a default zero value for expire time:")
	fmt.Printf("%#v\n", b2)

	var b3 Budget
	fmt.Println("Struct with default values")
	fmt.Printf("%#v\n", b3)
}
