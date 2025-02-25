package main

import "fmt"

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		t := array[i]
		for ii := i; ii >= 0; ii-- {
			if array[ii] > t {
				array[ii], array[i] = t, array[ii]
				i--
			}
		}
	}
	return array
}

func NonConstructibleChange(coins []int) int {
	temp := 1
	result := 1

	coins = InsertionSort(coins)

	for i, v := range coins {
		if i+1 <= len(coins)-1 {
			temp += v
			if coins[i+1] >= temp+1 {
				result = temp
				break
			} else {
				result = temp
			}
		} else if i == len(coins)-1 && len(coins) != 1 {
			result += v
		} else {
			if coins[i] < temp+1 {
				result = temp + 1
			}
		}
	}

	return result
}

func main() {
	// 1, 1, 2, 3, 5, 7, 22
	// 20
	// coins := []int{5, 7, 1, 1, 2, 3, 22}

	// 6
	// coins := []int{1, 1, 1, 1, 1}

	// 1
	// coins := []int{87}

	// 1
	// coins := []int{2}

	// 29
	// coins := []int{1, 2, 3, 4, 5, 6, 7}

	// 3
	// 1, 1, 4, 5, 6, 8, 9
	// coins := []int{6, 4, 5, 1, 1, 8, 9}

	// 2
	coins := []int{1}

	fmt.Println(NonConstructibleChange(coins))
}
