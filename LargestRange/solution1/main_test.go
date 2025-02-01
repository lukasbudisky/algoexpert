package main

import (
	"testing"
)

func TestLargestRange(t *testing.T) {
	array := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
	expected := []int{0, 7}

	result := LargestRange(array)
	if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected values: %v, got values: %v", expected, result)
	}
}
