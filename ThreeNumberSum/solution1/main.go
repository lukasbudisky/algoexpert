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
	for i := 0; i < len(sum); i++ {
		for ii := i + 1; ii < len(sum); ii++ {
			if sum[i][0] > sum[ii][0] {
				sum[i], sum[ii] = sum[ii], sum[i]
			} else if sum[i][0] == sum[ii][0] {
				if sum[i][1] > sum[ii][0] {
					sum[i], sum[ii] = sum[ii], sum[i]
				} else if sum[i][1] == sum[ii][1] {
					if sum[i][2] > sum[ii][2] {
						sum[i], sum[ii] = sum[ii], sum[i]
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
