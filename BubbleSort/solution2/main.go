package main

import "fmt"

func walk(array []int, l, r, len int, sorted bool) []int {
	if !sorted && l == len {
		sorted = true
		l = 0
		r = 1
	}
	if array[r] < array[l] {
		sorted = false
		array[l], array[r] = array[r], array[l]
	}
	l++
	r++
	if sorted && l == len {
		return array
	}
	return walk(array, l, r, len, sorted)
}

func BubbleSort(array []int) []int {
	if len(array) >= 2 {
		return walk(array, 0, 1, len(array)-1, false)
	}
	return array
}

func main() {
	a := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(BubbleSort(a))
}
