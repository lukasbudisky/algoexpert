package main

import "fmt"

func RightSmallerThan(array []int) []int {
	for i := 0; i < len(array); i++ {
		sum := 0
		for ii := i + 1; ii < len(array); ii++ {
			if array[i] > array[ii] {
				sum++
			}
		}
		array[i] = sum
	}
	return array
}

func main() {
	a := []int{8, 5, 11, -1, 3, 4, 2}

	fmt.Println(RightSmallerThan(a))
}
