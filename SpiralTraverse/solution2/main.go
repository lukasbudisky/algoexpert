package main

import "fmt"

func main() {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	// array := [][]int{
	// 	{1, 2, 3},
	// 	{12, 13, 4},
	// 	{11, 14, 5},
	// 	{10, 15, 6},
	// 	{9, 8, 7},
	// }
	// array := [][]int{
	// 	{1},
	// 	{3},
	// 	{2},
	// 	{5},
	// 	{4},
	// 	{7},
	// 	{6},
	// }
	// array := [][]int{
	// 	{1, 3, 2, 5, 4, 7, 6},
	// }
	// array := [][]int{
	// 	{1},
	// }

	fmt.Println(SpiralTraverse(array))
}

func SpiralTraverse(array [][]int) []int {
	var result []int
	sCp, sRp, cP, rP := 0, 0, 0, 0
	eCp := len(array[0]) - 1
	eRp := len(array) - 1

	for {
		result = append(result, array[rP][cP])
		if cP < eCp && rP <= eRp && rP == sRp {
			cP++
		} else if cP == eCp && rP < eRp && sRp != eRp {
			rP++
		} else if rP == eRp && cP > sCp && cP <= eCp && sRp != eRp {
			cP--
		} else if cP == sCp && rP <= eRp {
			if rP != sRp && sCp != eCp {
				rP--
				if rP == sRp {
					sCp++
					sRp++
					cP = sCp
					rP = sRp
					eCp--
					eRp--
					if sCp > eCp || sRp > eRp {
						break
					}
				}
			} else {
				break
			}
		} else if sCp >= eCp || sRp >= eRp {
			break
		}
	}

	return result
}
