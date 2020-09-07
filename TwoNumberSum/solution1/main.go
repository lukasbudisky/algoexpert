// Lukas Budisky

package main

import (
	"fmt"
	"time"
)

// TwoNumberSum function: O(n^2) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	for i, v := range array {
		for ii, vv := range array {
			if i > ii {
				if v+vv == target {
					return []int{v, vv}
				}
			}
		}
	}
	return []int{}
}

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	t := 10

	start := time.Now()
	fmt.Println(TwoNumberSum(a, t))
	elapsed := time.Since(start)
	fmt.Printf("Time: %s\n", elapsed)
}
