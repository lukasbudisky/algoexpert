package main

import (
	"fmt"
)

func FindThreeLargestNumbers(array []int) []int {
	r := make([]int, 3)
	r[0], r[1], r[2] = -1<<63, -1<<63, -1<<63
	for _, v := range array {
		if v >= r[2] {
			r[0] = r[1]
			r[1] = r[2]
			r[2] = v
		}
		if v < r[2] && v >= r[1] {
			r[0] = r[1]
			r[1] = v
		}
		if v < r[1] && v >= r[0] {
			r[0] = v
		}
	}
	return r
}

func main() {
	a := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	fmt.Println(FindThreeLargestNumbers(a))
}
