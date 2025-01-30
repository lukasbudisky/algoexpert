package main

import "fmt"

// str1[r-1] == str2[c-1] : E[r][c] = E[r-1][c-1] ? E[r][c] = 1 + min(E[r][c-1], E[r-1][c], E[r-1][c-1])

//		" "	 y	 a	 b	 d
// " "	 0	 1	 2	 3	 4
//  a	 1	 1	 1	 2	 3
//  b	 2	 2	 2	 1	 2
// 	c	 3	 3	 3   2   2
// result 2

func LevenshteinDistance(a, b string) int {
	a = " " + a
	b = " " + b

	array := make([][]int, len(a))
	for i := range array {
		array[i] = make([]int, len(b))
	}

	for i1, v1 := range a {
		for i2, v2 := range b {
			if v1 != v2 {
				if i1-1 < 0 || i2-1 < 0 {
					array[i1][i2] = i2 + i1
				} else {
					// value := min(array[i1][i2-1], array[i1-1][i2], array[i1-1][i2-1])
					if array[i1][i2-1] <= array[i1-1][i2] && array[i1][i2-1] <= array[i1-1][i2-1] {
						array[i1][i2] = 1 + array[i1][i2-1]
					} else if array[i1-1][i2] <= array[i1][i2-1] && array[i1-1][i2] <= array[i1-1][i2-1] {
						array[i1][i2] = 1 + array[i1-1][i2]
					} else if array[i1-1][i2-1] <= array[i1][i2-1] && array[i1-1][i2-1] <= array[i1-1][i2] {
						array[i1][i2] = 1 + array[i1-1][i2-1]
					}
				}
			} else {
				if i1 == 0 && i2 == 0 {
					array[i1][i2] = 0
				} else {
					array[i1][i2] = array[i1-1][i2-1]
				}
			}
		}
	}

	// fmt.Println(array)
	return array[len(a)-1][len(b)-1]
}

func main() {
	// 2
	str1 := "abc"
	str2 := "yabd"
	//
	// str1 := "algoexpert"
	// str2 := "algozexpert"
	// str1 := ""
	// str2 := "abc"
	// 4
	// str1 := "biting"
	// str2 := "mitten"
	fmt.Println(LevenshteinDistance(str1, str2))
}
