package main

import "fmt"

func increase(x, y int) bool {
	if x < y {
		return true
	}
	return false
}

func IsMonotonic(array []int) bool {
	i := 0
	mon := false

	if len(array) == 0 || len(array) == 1 {
		return true
	}
	inc := increase(array[0], array[len(array)-1])
	for {
		if i+1 == len(array) {
			break
		}
		if inc && array[i] <= array[i+1] {
			i++
			mon = true
		} else if !inc && array[i] >= array[i+1] {
			i++
			mon = true
		} else {
			mon = false
			break
		}
	}
	return mon
}

func main() {
	// a := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	a := []int{1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 8, 9, 10, 11}

	fmt.Println(IsMonotonic(a))
}
