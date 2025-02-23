package main

import "fmt"

func FirstNonRepeatingCharacter(str string) int {
	temp := make(map[rune]int)
	result := -1

	for i, v := range str {
		_, ok := temp[v]
		if !ok {
			temp[v] = i
		} else {
			temp[v] = -1
		}
	}

	// for k, v := range temp {
	// 	fmt.Printf("%s - %d\n", string(k), v)
	// }

	for _, v := range temp {
		if result == -1 {
			result = v
		} else if v <= result && v != -1 {
			result = v
		}
	}

	return result
}

func main() {
	// 6
	s := "faadabcbbebdf"
	// 1
	// s := "abcdcaf"

	fmt.Println(FirstNonRepeatingCharacter(s))
}
