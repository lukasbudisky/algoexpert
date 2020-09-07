// Lukas Budisky

package main

import (
	"fmt"
	"sort"
	"time"
)

// TwoNumberSum function: O(nlog(n)) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	lp, rp := 0, len(array)-1
	sort.Ints(array)
	for lp < rp {
		r := array[lp] + array[rp]
		switch {
		case r < target:
			lp++
		case r > target:
			rp--
		case r == target:
			return []int{array[lp], array[rp]}
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
