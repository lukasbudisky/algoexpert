package main

import (
	"fmt"
	"time"
)

// IsValidSubsequence function:
func IsValidSubsequence(array []int, sequence []int) bool {
	lp := 0
	if len(sequence) <= len(array) {
		for _, v := range array {
			if lp == len(sequence) {
				break
			}
			if v == sequence[lp] {
				lp++
			}
		}
	}
	return lp == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	start := time.Now()
	fmt.Println(IsValidSubsequence(array, sequence))
	elapsed := time.Since(start)
	fmt.Printf("Time: %s", elapsed)
}
