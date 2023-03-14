package main

import "fmt"

func IsPalindrome(str string) bool {
	l := 0
	r := len(str) - 1
	for {
		if l == r || l > r {
			break
		}
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	s := "paap"

	fmt.Println(IsPalindrome(s))
}
