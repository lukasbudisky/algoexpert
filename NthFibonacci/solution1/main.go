package main

import "fmt"

func fb(x, y, n int) int {
	if n == 1 || n == 0 {
		return x
	}
	return fb(y, x+y, n-1)
}

func GetNthFib(n int) int {
	return fb(0, 1, n)
}

func main() {
	fmt.Println(GetNthFib(0))
}
