package main

import (
	"fmt"
	"sort"
)

func FourNumberSum(array []int, target int) [][]int {
	var sum [][]int

	sort.Ints(array)
	for i := 0; i < len(array); i++ {
		for ii := i + 1; ii < len(array); ii++ {
			l := ii + 1
			r := len(array) - 1
			for {
				if l >= r {
					break
				}
				if array[i]+array[ii]+array[l]+array[r] == target {
					sum = append(sum, []int{array[r], array[l], array[ii], array[i]})
					l++
					r--
				} else if array[i]+array[ii]+array[l]+array[r] > target {
					r--
				} else if array[i]+array[ii]+array[l]+array[r] < target {
					l++
				}
			}
		}
	}
	return sum
}

func main() {
	// a := []int{7, 6, 4, -1, 1, 2}
	a := []int{-2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// t := 16
	t := 4

	fmt.Println(FourNumberSum(a, t))
}
