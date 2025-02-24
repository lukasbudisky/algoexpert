package main

import "fmt"

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	winners := make(map[string]int)
	result := ""

	for i, comp := range competitions {
		if results[i] == HOME_TEAM_WON {
			winners[comp[1]] += 0
			winners[comp[0]] += 3
		} else {
			winners[comp[1]] += 3
			winners[comp[0]] += 0
		}
	}

	temp := 0
	for k, v := range winners {
		if temp < v {
			result = k
			temp = v
		}
	}

	return result
}

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	fmt.Println(TournamentWinner(competitions, results))

}
