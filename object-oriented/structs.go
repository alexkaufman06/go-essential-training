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

func NewBudget(campaignId string, balance float64, expires time.Time) (*Budget, error) {
	if campaignId == "" {
		return nil, fmt.Errorf("empty campaignId")
	}

	if balance < 0 {
		return nil, fmt.Errorf("balance must be greater than 0")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("expiration date must be in the future")
	}

	b := Budget{
		CampaignId: campaignId,
		Balance:    balance,
		Expires:    expires,
	}
	return &b, nil // return pointer to budget value
}

func (b Budget) TimeLeft() time.Duration { // (b Budget) is a receiver, it adds a method to the Budget struct
	return b.Expires.Sub(time.Now().UTC())
}

func (b *Budget) Update(sum float64) { // Because we are changing the balance within the struct, we pass a pointer *
	b.Balance += sum
}

func main() {
	b1 := Budget{
		CampaignId: "Give rich people more money",
		Balance:    1000000,
		Expires:    time.Now().UTC().Add(7 * 24 * time.Hour),
	}

	fmt.Println("Print the contents of the struct:")
	fmt.Println(b1)
	fmt.Println("Print the time left for the above Budget:")
	fmt.Println(b1.TimeLeft())
	fmt.Println("Print the balance after subtracting 100000:")
	b1.Update(-100)
	fmt.Println(b1.Balance)
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

	expires := time.Now().UTC().Add(10 * 24 * time.Hour)
	b4, err := NewBudget("money for schools", 100.5, expires)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", b4)
	}

	b5, err := NewBudget("money for fools", -100.5, expires)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", b5)
	}

}
