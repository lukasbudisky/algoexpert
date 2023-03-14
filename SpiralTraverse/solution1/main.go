package main

import "fmt"

// Doesn't work for all tests
func SpiralTraverse(array [][]int) []int {
	var result []int
	r, rowStart := 0, 0
	rowEnd := len(array) - 1
	c, columnStart := 0, 0
	columnEnd := len(array[rowEnd]) - 1
	for {
		if r == rowStart {
			for i := columnStart; i <= columnEnd; i++ {
				result = append(result, array[r][i])
			}
			r++
			c = columnEnd
		} else if r == rowEnd {
			for i := columnEnd; i >= columnStart; i-- {
				result = append(result, array[r][i])
			}
			r--
			if c == columnEnd && r == rowStart {
				columnStart++
				columnEnd--
				rowStart++
				rowEnd--
				r = rowStart
			} else {
				c = columnStart
			}
		} else if c == columnStart {
			result = append(result, array[r][c])
			r--
			if r == rowStart {
				columnStart++
				columnEnd--
				rowStart++
				rowEnd--
				c = columnStart
				r = rowStart
			}
		} else if c == columnEnd {
			result = append(result, array[r][c])
			r++
		}
		if c > columnEnd || r > rowEnd {
			break
		}
	}

	return result
}

func main() {
	// a := [][]int{
	// 	{1, 2, 3, 4, 100},
	// 	{12, 13, 14, 5, 101},
	// 	{11, 16, 15, 6, 102},
	// 	{10, 9, 8, 7, 103},
	// }
	// a := [][]int{
	// 	{1, 2, 3, 4},
	// 	{10, 11, 12, 5},
	// 	{9, 8, 7, 6},
	// }
	// a := [][]int{
	// 	{1, 2, 3, 4},
	// 	{12, 13, 14, 5},
	// 	{11, 16, 15, 6},
	// 	{10, 9, 8, 7},
	// 	{101, 102, 103, 104},
	// }
	// a := [][]int{
	// 	{1, 2, 3, 4},
	// 	{12, 13, 14, 5},
	// 	{11, 16, 15, 6},
	// 	{10, 9, 8, 7},
	// }
	// a := [][]int{
	// 	{1, 2, 3},
	// 	{8, 9, 5},
	// 	{7, 6, 5},
	// }
	a := [][]int{
		{1, 2, 3},
		{12, 13, 4},
		{11, 14, 5},
		{10, 15, 6},
		{9, 8, 7},
	}
	// a := [][]int{
	// 	{1, 2, 3, 4},
	// 	{8, 7, 6, 5},
	// }
	// a := [][]int{
	// 	{1, 2},
	// 	{8, 3},
	// 	{7, 4},
	// 	{6, 5},
	// }
	// a := [][]int{
	// 	{1},
	// }

	fmt.Println(SpiralTraverse(a))
}
