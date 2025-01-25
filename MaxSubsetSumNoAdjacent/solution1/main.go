package main

import "fmt"

func MaxSubsetSumNoAdjacent(array []int) int {
	temp := []int{}
	if len(array) == 0 {
		return 0
	}
	for i, value := range array {
		if i == 0 {
			temp = append(temp, value)
		} else if i == 1 {
			if temp[0] > value {
				temp = append(temp, temp[0])
			} else {
				temp = append(temp, value)
			}
		} else {
			if temp[i-1] > (temp[i-2] + array[i]) {
				temp = append(temp, temp[i-1])
			} else {
				temp = append(temp, temp[i-2]+array[i])
			}
		}

	}

	return temp[len(temp)-1]
}

// maSums[i-1] maxSums[i-2]+array[i]
func main() {
	array := []int{75, 105, 120, 75, 90, 135}
	fmt.Println(MaxSubsetSumNoAdjacent(array))
}
