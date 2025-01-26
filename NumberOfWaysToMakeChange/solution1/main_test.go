package main

import "testing"

func TestNumberOfWaysToMakeChange1(t *testing.T) {
	n := 6
	denoms := []int{1, 5}
	result := 2

	test := NumberOfWaysToMakeChange(n, denoms)
	if test != result {
		t.Errorf("expected value: %d, got value from test: %d", result, test)
	}

}

func TestNumberOfWaysToMakeChange2(t *testing.T) {
	n := 10
	denoms := []int{1, 5, 10, 25}
	result := 4

	test := NumberOfWaysToMakeChange(n, denoms)
	if test != result {
		t.Errorf("expected value: %d, got value from test: %d", result, test)
	}
}

func TestNumberOfWaysToMakeChange3(t *testing.T) {
	n := 7
	denoms := []int{2, 3, 4, 7}
	result := 3

	test := NumberOfWaysToMakeChange(n, denoms)
	if test != result {
		t.Errorf("expected value: %d, got value from test: %d", result, test)
	}
}

func TestNumberOfWaysToMakeChange4(t *testing.T) {
	n := 0
	denoms := []int{5}
	result := 1

	test := NumberOfWaysToMakeChange(n, denoms)
	if test != result {
		t.Errorf("expected value: %d, got value from test: %d", result, test)
	}
}
