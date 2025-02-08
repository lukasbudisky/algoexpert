package main

import "fmt"

func ZigzagTraverse(array [][]int) []int {
	result := []int{}
	down := true
	x, y := 0, 0
	borUp := 0
	borDown := len(array) - 1
	borRight := len(array[0]) - 1
	borLeft := 0
	result = append(result, array[y][x])

	for {
		if x > len(array[0])-1 || y > len(array)-1 {
			break
		}

		if down {
			if y == borUp && x == borLeft {
				if y+1 <= borDown {
					y++
					result = append(result, array[y][x])
				}
				down = false
			} else if y == borUp {
				y++
				if x-1 >= borLeft {
					x--
					if y <= borDown {
						result = append(result, array[y][x])
					} else {
						// test
						y--
						x++
						down = false
					}
				}
			} else if x == borLeft || y == borDown {
				if y+1 <= borDown {
					y++
					result = append(result, array[y][x])
				} else {
					if x+1 < borRight {
						x++
						result = append(result, array[y][x])
					}
				}
				down = false
			} else {
				y++
				x--
				result = append(result, array[y][x])
			}
		} else {
			if x+1 <= borRight && y < borDown {
				x++
				if y-1 >= borUp {
					y--
				}
				result = append(result, array[y][x])
			} else if y <= borDown && x+1 > borRight {
				y++
				if y <= borDown {
					result = append(result, array[y][x])
				}
				down = true
			} else {
				x++
				if x == borRight && y == borDown {
					if len(array)%2 == 1 && len(array[y])%2 != 1 {
						result = append(result, array[y][x])
					} else if len(array)%2 == 0 && len(array[y])%2 == 1 {
						result = append(result, array[y][x])
					} else {
						result = append(result, array[y-1][x])
						result = append(result, array[y][x])
					}
					// result = append(result, array[y][x])
					break
				} else if y-1 >= borUp {
					y--
					result = append(result, array[y][x])
				} else {
					result = append(result, array[y][x])
				}
			}
			if y == borUp {
				if x+1 <= borRight {
					x++
					result = append(result, array[y][x])
				} else {
					if y+1 <= borDown {
						y++
						if x <= borRight {
							result = append(result, array[y][x])
						}
					}
				}
				down = true
			}

		}
	}

	return result
}

func main() {
	// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
	array := [][]int{
		{1, 3},
		{2, 4},
		{5, 7},
		{6, 8},
		{9, 10},
	}
	// array := [][]int{
	// 	{1, 3, 4, 10},
	// 	{2, 5, 9, 11},
	// 	{6, 8, 12, 15},
	// 	{7, 13, 14, 16},
	// }
	// array := [][]int{
	// 	{1, 3, 4, 7, 8},
	// 	{2, 5, 6, 9, 10},
	// }
	// array := [][]int{
	// 	{1, 3, 4, 10, 11},
	// 	{2, 5, 9, 12, 19},
	// 	{6, 8, 13, 18, 20},
	// 	{7, 14, 17, 21, 24},
	// 	{15, 16, 22, 23, 25},
	// }
	// array := [][]int{
	// 	{1},
	// 	{2},
	// 	{3},
	// 	{4},
	// 	{5},
	// }
	// array := [][]int{
	// 	{1, 2, 3, 4, 5},
	// }
	// array := [][]int{
	// 	{1},
	// }
	fmt.Println(ZigzagTraverse(array))
}
