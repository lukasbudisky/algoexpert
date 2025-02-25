package main

import "fmt"

func TransposeMatrix(matrix [][]int) [][]int {
	x, y := 0, 0
	result := [][]int{}

	for {
		temp := []int{}
		if x == len(matrix[y]) {
			break
		}
		for {
			if y == len(matrix) {
				break
			}
			temp = append(temp, matrix[y][x])
			y++
		}
		result = append(result, temp)
		x++
		y = 0
	}

	return result
}

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(TransposeMatrix(input))
}
