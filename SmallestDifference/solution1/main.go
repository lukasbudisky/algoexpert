package main

import (
	"fmt"
	"math"
)

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SmallestDifference(array1, array2 []int) []int {
	max := math.MaxInt32
	var r []int

	for i := 0; i < len(array1); i++ {
		for ii := 0; ii < len(array2); ii++ {
			sum := absolute(array1[i] - array2[ii])
			if sum < max {
				r = []int{array1[i], array2[ii]}
				max = sum
			}
		}
	}
	return r
}

func main() {
	a1 := []int{-1, 5, 10, 20, 28, 3}
	a2 := []int{26, 134, 135, 15, 17}

	fmt.Println(SmallestDifference(a1, a2))
}
