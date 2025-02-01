package main

import (
	"fmt"
)

func LargestRange(array []int) []int {
	result := []int{}

	m := make(map[int]bool)

	for _, v := range array {
		m[v] = true
	}
	// fmt.Println(m)
	for _, v := range array {
		temp := []int{}
		// fmt.Println(v)
		if m[v] {
			start := v
			end := v
			for {
				if m[start] {
					start--
				} else {
					break
				}
			}
			for {
				if m[end] {
					end++
				} else {
					break
				}
			}
			for i := start + 1; i < end; i++ {
				temp = append(temp, i)
			}
		}
		if len(temp) > len(result) {
			result = temp
		}
		// fmt.Println(temp)
	}

	return []int{result[0], result[len(result)-1]}
}

func main() {
	array := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
	fmt.Println(LargestRange(array))
}
