package main

import "fmt"

func GetNthFib(n int) int {
	x := 0
	y := 1
	if n == 1 {
		y = 0
	}
	for i := 0; i < n-2; i++ {
		t := x + y
		x = y
		y = t
	}
	return y
}

func main() {
	fmt.Println(GetNthFib(6))
}
