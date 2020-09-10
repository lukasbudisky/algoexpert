package main

import (
	"fmt"
	"time"
)

// IsValidSubsequence function:
func IsValidSubsequence(array []int, sequence []int) bool {
	iS := 0
	iA := 0
	for iA < len(array) && iS < len(sequence) {
		if array[iA] == sequence[iS] {
			iS++
		}
		iA++
	}
	return iS == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	start := time.Now()
	fmt.Println(IsValidSubsequence(array, sequence))
	elapsed := time.Since(start)
	fmt.Printf("Time: %s", elapsed)
}
