package main

import "fmt"

func walk(array []int, l, h, target int) int {
	m := (l + h) / 2
	if l > h {
		return -1
	} else if target > array[m] {
		return walk(array, m+1, h, target)
	} else if target < array[m] {
		return walk(array, l, m-1, target)
	} else if target == array[m] {
		return m
	} else {
		return -1
	}
}

func BinarySearch(array []int, target int) int {
	return walk(array, 0, len(array)-1, target)
}

func main() {
	a := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	// b := []int{1, 5, 23, 111}

	// fmt.Println(BinarySearch(b, 35))
	// fmt.Println(BinarySearch(a, 0))
	// fmt.Println(BinarySearch(a, 1))
	// fmt.Println(BinarySearch(a, 21))
	fmt.Println(BinarySearch(a, 100))
	// fmt.Println(BinarySearch(a, 45))
	// fmt.Println(BinarySearch(a, 61))
	// fmt.Println(BinarySearch(a, 71))
	// fmt.Println(BinarySearch(a, 72))
	// fmt.Println(BinarySearch(a, 73))
	// fmt.Println(BinarySearch(a, 100))
}
