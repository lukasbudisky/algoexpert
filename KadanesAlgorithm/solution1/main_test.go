package main

import "testing"

func TestKadanesAlgorithm1(t *testing.T) {
	array := []int{-2, 1}
	expected := 1

	test := KadanesAlgorithm(array)
	if test != expected {
		t.Errorf("expected output: %d, got output: %d", expected, test)
	}
}

func TestKadanesAlgorithm2(t *testing.T) {
	array := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	expected := 19

	test := KadanesAlgorithm(array)
	if test != expected {
		t.Errorf("expected output: %d, got output: %d", expected, test)
	}
}
