package main

import "fmt"

func BubbleSort(array []int) []int {
	sorted := false
	length := len(array)
	l := 0
	r := 1
	if length >= 2 {
		for {
			if !sorted && l == length-1 {
				sorted = true
				l = 0
				r = 1
			}
			if array[r] < array[l] {
				sorted = false
				t := array[l]
				array[l] = array[r]
				array[r] = t
			}
			l++
			r++
			if sorted && l == length-1 {
				break
			}
		}
	}
	return array
}

func main() {
	a := []int{8, 5, 2, 9, 5, 6, 3}

	fmt.Println(BubbleSort(a))
}
