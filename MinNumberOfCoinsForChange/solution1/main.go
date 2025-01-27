package main

import (
	"fmt"
)

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	nums := make([]int, n+1)
	x := 9999

	for i := range nums {
		nums[i] = x
	}
	nums[0] = 0

	for _, d := range denoms {
		for i := range nums {
			if d <= i {
				if nums[i] >= 1+nums[i-d] {
					nums[i] = 1 + nums[i-d]
				}
			}
		}
		// fmt.Println(nums)
	}
	if nums[n] == x {
		return -1
	} else {
		return nums[n]
	}
}

func main() {
	// nums[i] >= 1 + nums[i-d]
	//   | 0 1 2 3 4 5 6 7
	// --------------------
	// 2 | 0 0 1 1 2 2 3 3
	// 4 | 0 0 1 1 1 1 1 2
	n := 7
	denoms := []int{2, 4}
	fmt.Println(MinNumberOfCoinsForChange(n, denoms))
}
