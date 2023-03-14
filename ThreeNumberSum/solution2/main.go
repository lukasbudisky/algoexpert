package main

import (
	"fmt"
	"sort"
)

// ThreeNumberSum function: O(n^3) time
func ThreeNumberSum(array []int, target int) [][]int {
	var sum [][]int
	sort.Ints(array)
	for i := 0; i < len(array); i++ {
		l := i + 1
		r := len(array) - 1
		for {
			if l >= r {
				break
			}
			if array[i]+array[l]+array[r] == target {
				var t []int
				t = append(t, array[i], array[l], array[r])
				sum = append(sum, t)
				r--
				l++
			} else if array[i]+array[l]+array[r] < target {
				l++
			} else if array[i]+array[l]+array[r] > target {
				r--
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
