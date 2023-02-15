package main

import (
	"encoding/csv"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

type challengeTestCase struct {
	value    string
	expected string
}

// Below is my solution, but the video solution created a utility function for downloading the csv:
// https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch07/07_05/sqrt_test.go

func TestChallenge(t *testing.T) {
	csvFile, err := os.Open("test_cases.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("successfully opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		tc := challengeTestCase{
			value:    line[0],
			expected: line[1],
		}

		val, err := strconv.ParseFloat(tc.value, 64)
		if err != nil {
			t.Fatalf("error parsing value: %s", err)
		}

		expected, err := strconv.ParseFloat(tc.expected, 64)
		if err != nil {
			t.Fatalf("error parsing expected: %s", err)
		}

		resp, err := Sqrt(val)
		require.NoError(t, err)
		require.InDelta(t, expected, resp, 0.001)
	}
}
