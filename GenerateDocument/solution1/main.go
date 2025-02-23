package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	charMap := make(map[rune]int)
	docMap := make(map[rune]int)
	result := true

	for _, v := range characters {
		charMap[v]++
	}

	for _, v := range document {
		docMap[v]++
	}

	for k, v := range docMap {
		_, ok := charMap[k]
		if ok {
			if charMap[k] >= v {
				continue
			} else {
				result = false
			}
		} else {
			result = false
		}
	}

	// for k, v := range charMap {
	// 	fmt.Printf("char %s - %d\n", string(k), v)
	// }

	// fmt.Println("-----")

	// for k, v := range docMap {
	// 	fmt.Printf("document %s - %d\n", string(k), v)
	// }

	return result
}

func main() {
	// characters := "Bste!hetsi ogEAxpelrt x "
	characters := "Bste!hetsi ogEAxpert"
	document := "AlgoExpert is the Best!"
	fmt.Println(GenerateDocument(characters, document))
}
