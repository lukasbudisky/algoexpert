package main

import "fmt"

func RunLengthEncoding(str string) string {
	result := ""
	count := 1

	for i, v := range []byte(str) {
		if i+1 <= len(str)-1 {
			if count < 9 {
				if v == str[i+1] {
					count++
				} else {
					result = fmt.Sprintf("%s%d%s", string(result), count, string(v))
					count = 1
				}
			} else {
				result = fmt.Sprintf("%s%d%s", string(result), count, string(v))
				count = 1
			}
		} else {
			result = fmt.Sprintf("%s%d%s", string(result), count, string(v))
		}
	}

	return result
}

func main() {
	// s := "AAAAAAAAAAAAABBCCCCDD"
	s := "........______=========AAAA   AAABBBB   BBB"
	fmt.Println(RunLengthEncoding(s))
}
