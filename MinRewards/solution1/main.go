package main

import "fmt"

func MinRewards(scores []int) int {
	temp := make([]int, len(scores))
	result := 0
	for i := range temp {
		temp[i] = 1
	}
	for i := 1; i <= len(scores)-1; i++ {
		if scores[i-1] < scores[i] {
			temp[i] = temp[i-1] + 1
		}
	}
	// fmt.Println(temp)
	for i := len(scores) - 1; i > 0; i-- {
		if scores[i-1] > scores[i] {
			if temp[i-1] > temp[i]+1 {
				continue
			} else {
				temp[i-1] = temp[i] + 1
			}
		}
	}
	// fmt.Println(temp)
	for _, v := range temp {
		result += v
	}
	return result
}

func main() {
	// 25
	array := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	// 3 [2,1]
	// array := []int{10, 5}
	// 93
	// array := []int{800, 400, 20, 10, 30, 61, 70, 90, 17, 21, 22, 13, 12, 11, 8, 4, 2, 1, 3, 6, 7, 9, 0, 68, 55, 67, 57, 60, 51, 661, 50, 65, 53}
	fmt.Println(MinRewards(array))
}
