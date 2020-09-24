package main

import "fmt"

func GetNthFib(n int) int {
	r := 0
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		r = x
		t := x + y
		x = y
		y = t
	}
	return r
}

func main() {
	fmt.Println(GetNthFib(6))
}
