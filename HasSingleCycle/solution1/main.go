package main

import "fmt"

func HasSingleCycle(array []int) bool {
	temp := make(map[int]bool)
	i := 0
	c := 0
	for {
		if c == len(array) {
			break
		}
		v := array[i] % len(array)
		switch {
		case i+v > 0 && i+v <= len(array)-1:
			p := i + v
			temp[p] = true
			i = p
		case i+v > 0 && i+v >= len(array)-1:
			p := i + v - len(array)
			temp[p] = true
			i = p
		case i+v < 0 && i+v <= len(array)-1:
			p := len(array) + i + v
			temp[p] = true
			i = p
		case i+v == 0:
			p := 0
			temp[p] = true
			i = p
		}
		c++
	}

	for i := range array {
		v, ok := temp[i]
		if !ok || !v {
			return false
		}
	}

	return true
}

func main() {
	// array := []int{2, 3, 1, -4, -4, 2}
	// array := []int{0, 1, 1, 1, 1}
	// array := []int{1, -1, 1, -1}
	array := []int{10, 11, -6, -23, -2, 3, 88, 908, -26}
	fmt.Println(HasSingleCycle(array))
}
