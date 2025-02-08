package main

import "fmt"

func LongestSubarrayWithSum(array []int, targetSum int) []int {
	result := []int{}
	for i1, v1 := range array {
		temp := v1
		if temp == targetSum && len(result) != 2 {
			result = []int{i1, i1}
		} else if i1+1 <= len(array)-1 && len(array) != 1 {
			for i2, v2 := range array[i1+1:] {
				temp += v2
				index := i2 + len(array[0:i1]) + 1
				if temp == targetSum {
					if len(result) == 2 {
						if index-i1 > result[1]-result[0] {
							result = []int{i1, index}
						}
					} else {
						result = []int{i1, index}
					}
				} else if temp > targetSum {
					break
				}
			}
		} else if len(array) == 1 && targetSum == temp {
			result = []int{i1, i1}
		}
	}

	return result
}

func main() {
	// [4,8]
	// array := []int{1, 2, 3, 4, 3, 3, 1, 2, 1}
	// targetSum := 10
	// [0,9]
	// array := []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}
	// targetSum := 1
	// [0,0]
	// array := []int{10}
	// targetSum := 10
	// [8,8]
	array := []int{1, 4, 10, 15, 31, 7, 1, 40, 0, 20, 1, 1, 1, 1, 2, 1}
	targetSum := 0
	fmt.Println(LongestSubarrayWithSum(array, targetSum))
}
