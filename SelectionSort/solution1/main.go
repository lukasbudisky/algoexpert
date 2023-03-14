package main

import "fmt"

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		si := i
		for ii := i + 1; ii < len(array); ii++ {
			if array[si] > array[ii] {
				si = ii
			}
		}
		array[i], array[si] = array[si], array[i]
	}

	return array
}

func main() {
	a := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(SelectionSort(a))
}
