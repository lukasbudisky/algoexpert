package main

import "fmt"

type SpecialArray []interface{}

func walk(array SpecialArray, n int) int {
	sum := 0
	n++
	for _, v := range array {
		if vv, ok := v.(SpecialArray); ok {
			sum += walk(vv, n)
		} else {
			sum += v.(int)
		}
	}
	return sum * n
}

func ProductSum(array []interface{}) int {
	return walk(array, 0)
}

func main() {
	a := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{-13, 8},
			4,
		},
	}
	fmt.Println(a)
	fmt.Println(ProductSum(a))
}
