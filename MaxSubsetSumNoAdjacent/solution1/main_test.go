package main

import "testing"

func TestMaxSubsetSumNoAdjacentArray(t *testing.T) {
	array := []int{75, 105, 120, 75, 90, 135}
	expected := 330

	result := MaxSubsetSumNoAdjacent(array)
	if expected != result {
		t.Errorf("values are not equal. Expected value %d, got value: %d", expected, result)
	}
}

func TestMaxSubsetSumNoAdjacentEmptyArray(t *testing.T) {
	array := []int{}
	expected := 0

	result := MaxSubsetSumNoAdjacent(array)
	if expected != result {
		t.Errorf("values are not equal. Expected value %d, got value: %d", expected, result)
	}
}
