package main

import "fmt"

func MoveElementToEnd(array []int, toMove int) []int {
	p := 0
	for i := 0; i < len(array); i++ {
		if array[i] != toMove {
			array[p], array[i] = array[i], array[p]
			p++
		}
	}
	return array
}

func main() {
	a := []int{2, 1, 2, 2, 2, 3, 4, 2}
	t := 2

	fmt.Println(MoveElementToEnd(a, t))
}
