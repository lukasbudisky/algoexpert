package main

import "fmt"

func CommonCharacters(strings []string) []string {
	result := []string{}
	resultMap := make(map[rune]int)

	for _, s := range strings {
		temp := make(map[rune]int)
		for _, c := range s {
			temp[c]++
		}
		for k := range temp {
			resultMap[k]++
		}
	}

	for k, v := range resultMap {
		fmt.Printf("%s - %d\n", string(k), v)
		if v == len(strings) {
			result = append(result, string(k))
		}
	}
	return result
}

func main() {
	// s := []string{"abc", "bcd", "cbad"}
	s := []string{"ab&cdef!", "f!ed&cba", "a&bce!d", "ae&fb!cd", "efa&!dbc", "eff!&fff&fffffffbcda", "eeee!efff&fffbbbbbaaaaaccccdddd", "*******!***&****abdcef************", "*******!***&****a***********f*", "*******!***&****b***********c*"}
	fmt.Println(CommonCharacters(s))
}
