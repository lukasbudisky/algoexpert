package main

import "fmt"

func getSubArrayLow(array []int, root int) []int {
	subArray := []int{}
	for i, v := range array {
		if v < root && i != 0 {
			subArray = append(subArray, v)
		}
	}
	return subArray
}

func getSubArrayHigh(array []int, root int) []int {
	subArray := []int{}
	for i, v := range array {
		if v >= root && i != 0 {
			subArray = append(subArray, v)
		}
	}
	return subArray
}

func SameBsts(arrayOne, arrayTwo []int) bool {
	result := true
	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
		return true
	} else if len(arrayOne) != len(arrayTwo) || arrayOne[0] != arrayTwo[0] {
		return false
	} else {
		root := arrayOne[0]

		subArrayOneLow := getSubArrayLow(arrayOne, root)
		subArrayTwoLow := getSubArrayLow(arrayTwo, root)
		subArrayOneHigh := getSubArrayHigh(arrayOne, root)
		subArrayTwoHigh := getSubArrayHigh(arrayTwo, root)

		if !SameBsts(subArrayOneLow, subArrayTwoLow) {
			return false
		}
		if !SameBsts(subArrayOneHigh, subArrayTwoHigh) {
			return false
		}
	}

	return result
}

func main() {
	arrayOne := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
	arrayTwo := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}

	fmt.Println(SameBsts(arrayOne, arrayTwo))
}
