package main

import (
	"fmt"
	"sort"
)

// ThreeNumberSum function: O(n^3) time
func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	var sum [][]int = [][]int{}
	for i, v := range array {
		for ii, vv := range array {
			if i > ii {
				for iii, vvv := range array {
					if ii > iii {
						tSum := []int{}
						if (v + vv + vvv) == target {
							tSum = append(tSum, v, vv, vvv)
							sort.Ints(tSum)
							sum = append(sum, tSum)
						}
					}
				}
			}
		}
	}

	return sum
}

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	targetSum := 0

	fmt.Println(ThreeNumberSum(array, targetSum))
}
