package main

import (
	"testing"
)

func TestMinRewards1(t *testing.T) {
	array := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	expected := 25

	result := MinRewards(array)
	if expected != result {
		t.Errorf("expected result: %d, got result: %d", expected, result)
	}
}

func TestMinRewards2(t *testing.T) {
	array := []int{10, 5}
	expected := 3

	result := MinRewards(array)
	if expected != result {
		t.Errorf("expected result: %d, got result: %d", expected, result)
	}
}

func TestMinRewards3(t *testing.T) {
	array := []int{800, 400, 20, 10, 30, 61, 70, 90, 17, 21, 22, 13, 12, 11, 8, 4, 2, 1, 3, 6, 7, 9, 0, 68, 55, 67, 57, 60, 51, 661, 50, 65, 53}
	expected := 93

	result := MinRewards(array)
	if expected != result {
		t.Errorf("expected result: %d, got result: %d", expected, result)
	}
}
