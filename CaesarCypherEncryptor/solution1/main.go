package main

import "fmt"

const (
	min = 97
	max = 122
)

func CaesarCipherEncryptor(str string, key int) string {
	out := []rune(str)
	for i := 0; i < len(out); i++ {
		if int(out[i])+key > max {
			x := rune(int(out[i]) + key - max - 1)
			for {
				if x+min <= max {
					break
				}
				x = x + min - max - 1
			}
			out[i] = x + min
		} else {
			out[i] = rune(int(out[i]) + key)
		}
	}
	return string(out)
}

func main() {
	fmt.Println(CaesarCipherEncryptor("abc", 57))
}
