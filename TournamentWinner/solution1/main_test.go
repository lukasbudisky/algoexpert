package main

import "testing"

func TestTournamentWinner(t *testing.T) {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	expected := "Python"

	test := TournamentWinner(competitions, results)
	if test != expected {
		t.Errorf("expected output: %s, got output: %s", expected, test)
	}

}
