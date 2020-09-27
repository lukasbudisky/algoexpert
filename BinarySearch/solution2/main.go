package main

import "fmt"

func BinarySearch(array []int, target int) int {
	h := len(array) - 1
	for i := 0; i <= h; i++ {
		if h != 1 {
			if ((h - i%2) + i) == 1 {
				i = (h-i)/2 + 1 + i
			} else {
				i = (h-i)/2 + i
			}
		}
		if target < array[i] && h != i {
			h = i
			i = -1
		} else if target > array[i] && h != i {
			continue
		} else if target == array[i] {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}

	fmt.Println(BinarySearch(a, 21))
}
