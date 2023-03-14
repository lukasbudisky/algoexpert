package main

import "fmt"

func palindrome(str string, l, r int) bool {
	if l >= r {
		return true
	} else if str[l] == str[r] {
		return palindrome(str, l+1, r-1)
	}
	return false
}

func IsPalindrome(str string) bool {
	return palindrome(str, 0, len(str)-1)
}

func main() {
	s := "pabcap"

	fmt.Println(IsPalindrome(s))
}
