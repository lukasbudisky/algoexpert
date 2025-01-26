package main

import (
	"fmt"
)

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	ways := make([]int, n+1)
	ways[0] = 1

	// ways[10] += ways[10-5]
	// ways[amount] += ways[amount-denom]

	// d <= i]
	// ways[1] += ways[1-1]

	// ways[1] = ways[1] + ways[1-1]

	// ways[2] = ways[2] + ways[2-1]
	// ways[2] += ways[2-1]
	// ways[5] = ways[5] + ways[5-5]
	// 				0	 + 2
	for _, d := range denoms {
		for i, _ := range ways {
			if d <= i {
				ways[i] += ways[i-d]
			}
		}
		// fmt.Println(ways)
	}

	return ways[len(ways)-1]
}

func main() {
	n := 6
	denoms := []int{1, 5}
	// n := 10
	// denoms := []int{1, 5, 10, 25}
	// n := 7
	// denoms := []int{2, 3, 4, 7}
	// n := 0
	// denoms := []int{5}

	fmt.Println(NumberOfWaysToMakeChange(n, denoms))
}
