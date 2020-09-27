package main

import "fmt"

func BinarySearch(array []int, target int) int {
	l, m := 0, 0
	h := len(array) - 1

	for {
		m = (l + h) / 2
		if l > h {
			break
		}
		if target > array[m] {
			l = m + 1
		} else if target < array[m] {
			h = m - 1
		} else if target == array[m] {
			return m
		}
	}
	return -1
}

func main() {
	a := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	// b := []int{1, 5, 23, 111}

	// fmt.Println(BinarySearch(b, 35))
	// fmt.Println(BinarySearch(a, 0))
	// fmt.Println(BinarySearch(a, 1))
	// fmt.Println(BinarySearch(a, 21))
	fmt.Println(BinarySearch(a, 33))
	// fmt.Println(BinarySearch(a, 45))
	// fmt.Println(BinarySearch(a, 61))
	// fmt.Println(BinarySearch(a, 71))
	// fmt.Println(BinarySearch(a, 72))
	// fmt.Println(BinarySearch(a, 73))
	// fmt.Println(BinarySearch(a, 100))
}
