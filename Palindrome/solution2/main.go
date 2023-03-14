package main

import "fmt"

func IsPalindrome(str string) bool {
	var newStr []byte
	for i := len(str) - 1; i >= 0; i-- {
		newStr = append(newStr, str[i])
	}
	return str == string(newStr)
}

func main() {
	s := "abcdcba"

	fmt.Println(IsPalindrome(s))
}
