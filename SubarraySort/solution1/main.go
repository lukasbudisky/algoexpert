package main

import (
	"fmt"
)

func cmpValues(l, b, v int) (low int, big int) {
	if b < v {
		b = v
	}
	if l > v {
		l = v
	}
	return l, b
}

func SubarraySort(array []int) []int {
	b := -1
	l := 9999

	for i, v := range array {
		if i != 0 && i < len(array)-1 {
			if v < array[i-1] || v > array[i+1] {
				l, b = cmpValues(l, b, v)
			}
		} else {
			if i == 0 || i == len(array) {
				if v > array[i+1] && i != len(array)-1 {
					l, b = cmpValues(l, b, v)
				}
			} else {
				if v < array[i-1] && i > 0 {
					l, b = cmpValues(l, b, v)
				}
			}
		}
	}

	// fmt.Println(l, b)
	a1, a2 := -1, -1
	if b != -1 {
		for i, v := range array {
			if i < len(array)-1 {
				if l >= v && l < array[i+1] {
					a1 = i + 1
					break
				} else if l < v && l < array[i+1] {
					a1 = i
					break
				}
			} else {
				if i > 0 && l <= array[i-1] {
					a1 = i
				}
			}
		}
		for i, v := range array {
			if i > 0 && i < len(array)-1 {
				if b >= v && b < array[i+1] {
					a2 = i
					break
				} else {
					a2 = i + 1
				}
			} else {
				if i < len(array)-1 && b > array[i+1] {
					a2 = i
				}
			}
		}
	}

	if a1 < a2 {
		return []int{a1, a2}
	} else {
		return []int{a2, a1}
	}
}

func main() {
	// [3,9]
	// array := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	// [2,4]
	// array := []int{1, 2, 8, 4, 5}
	// [-1, -1]
	// array := []int{1, 2}
	// [-1, -1]
	// array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// [4, 24]
	// array := []int{1, 2, 3, 4, 5, 6, 18, 21, 22, 7, 14, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 4, 14, 11, 6, 33, 35, 41, 55}
	// [0,1]
	// array := []int{2, 1}
	// [0,12]
	array := []int{4, 8, 7, 12, 11, 9, -1, 3, 9, 16, -15, 51, 7}
	fmt.Println(SubarraySort(array))
}
