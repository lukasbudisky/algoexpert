package main

import "fmt"

func FindThreeLargestNumbers(array []int) []int {
	for i := 0; i < len(array); i++ {
		if array[i] > array[2] {
			t := array[2]
			array[2] = array[i]
			array[i] = t
		}
		if array[i] > array[1] && array[i] < array[2] {
			t := array[1]
			array[1] = array[i]
			array[i] = t
		}
		if array[i] > array[0] && array[i] < array[1] {
			t := array[0]
			array[0] = array[i]
			array[i] = t
		}
	}
	return array[0:3]
}

func main() {
	a := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	fmt.Println(FindThreeLargestNumbers(a))
}
