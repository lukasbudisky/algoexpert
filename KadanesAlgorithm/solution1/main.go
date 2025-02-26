package main

import "fmt"

func KadanesAlgorithm(array []int) int {
	result := array[0]
	for i, v := range array {
		temp := v
		if temp > result {
			result = temp
		}
		if i+1 <= len(array)-1 {
			for _, vv := range array[i+1:] {
				temp += vv
				if temp > result {
					result = temp
				}
			}
		}
	}

	return result
}

func main() {
	// 19
	// array := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}

	// 1
	array := []int{-2, 1}
	fmt.Println(KadanesAlgorithm(array))
}
