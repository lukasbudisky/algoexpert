// Lukas Budisky

package main

import (
	"fmt"
	"time"
)

// TwoNumberSum function: O(1) time | O(n) space
func TwoNumberSum(array []int, target int) []int {
	data := make(map[int]bool)
	for _, v := range array {
		y := target - v
		if _, ok := data[y]; ok {
			return []int{y, v}
		}
		data[v] = true
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
