package main

import "fmt"

func LongestPeak(array []int) int {
	result := 0
	position := 1
	peaks := []int{}

	for {
		if position >= (len(array) - 1) {
			break
		} else if array[position] > array[position+1] && array[position] > array[position-1] {
			peaks = append(peaks, position)
		}
		position++
	}

	for _, position := range peaks {
		left, right := position, position
		temp := 1

		for {
			if left-1 < 0 {
				break
			} else if array[left] > array[left-1] {
				temp++
				left--
			} else {
				break
			}
		}
		for {
			if right+1 > (len(array) - 1) {
				break
			} else if array[right] > array[right+1] {
				temp++
				right++
			} else {
				break
			}
		}
		if temp > result {
			result = temp
		}
	}
	return result
}

func main() {
	// array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	// array := []int{1, 1, 3, 2, 1}
	array := []int{1, 2, 3, 2, 1, 1}

	fmt.Println(LongestPeak(array))
}
