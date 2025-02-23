package main

import "fmt"

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		t := array[i]
		for ii := i; ii >= 0; ii-- {
			if array[ii] > t {
				array[ii], array[i] = t, array[ii]
				i--
			}
		}
	}
	return array
}

func SortedSquaredArray(array []int) []int {
	for i, v := range array {
		array[i] = v * v
	}
	return InsertionSort(array)
}

func main() {
	// array := []int{1, 2, 3, 5, 6, 8, 9}
	// array := []int{-3, -2, -1}
	array := []int{-10, -5, 0, 5, 10}

	fmt.Println(SortedSquaredArray(array))
}
