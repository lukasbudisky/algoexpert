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

func main() {
	a := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(InsertionSort(a))
}
