package main

import (
	"fmt"
	"time"
)

// IsValidSubsequence function:
func IsValidSubsequence(array []int, sequence []int) bool {
	dataSequence := map[int]int{}
	for i, v := range sequence {
		dataSequence[v] = i
	}
	dataArray := []int{}
	for _, v := range array {
		if _, ok := dataSequence[v]; ok {
			dataArray = append(dataArray, v)
		}
	}
	if len(sequence) <= len(dataArray) {
		for i, v := range sequence {
			if v != dataArray[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	start := time.Now()
	fmt.Println(IsValidSubsequence(array, sequence))
	elapsed := time.Since(start)
	fmt.Printf("Time: %s", elapsed)
}
