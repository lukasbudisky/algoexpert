package main

import "fmt"

func BinarySearch(array []int, target int) int {
	i := 0
	j := len(array) - 1
	for {
		if i == j {
			return -1
		}
		if array[i] < target {
			i++
		}
		if array[j] > target {
			j--
		}
		if array[i] == target {
			return i
		}
		if array[j] == target {
			return j
		}
	}
}

func main() {
	a := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}

	fmt.Println(BinarySearch(a, 33))
}
