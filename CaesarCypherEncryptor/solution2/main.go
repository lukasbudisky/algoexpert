package main

import "fmt"

const (
	min = 97
	max = 122
)

func CaesarCipherEncryptor(str string, key int) string {
	var newStr []rune
	modulo := max - min + 1
	for i := 0; i < len(str); i++ {
		if int(str[i])+key > max {
			newStr = append(newStr, rune((int(str[i])+key-max)%modulo+min-1))
		} else {
			newStr = append(newStr, rune(int(str[i])+key))
		}
	}
	return string(newStr)
}

func main() {
	fmt.Println(CaesarCipherEncryptor("xyz", 25))
}
