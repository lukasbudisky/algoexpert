package main

import "testing"

func TestHasSingleCycle1(t *testing.T) {
	expected := true
	array := []int{2, 3, 1, -4, -4, 2}
	test := HasSingleCycle(array)

	if test != expected {
		t.Errorf("expected value: %v, got value: %v", expected, test)
	}
}

func TestHasSingleCycle2(t *testing.T) {
	expected := false
	array := []int{0, 1, 1, 1, 1}
	test := HasSingleCycle(array)

	if test != expected {
		t.Errorf("expected value: %v, got value: %v", expected, test)
	}
}

func TestHasSingleCycle3(t *testing.T) {
	expected := false
	array := []int{1, -1, 1, -1}
	test := HasSingleCycle(array)

	if test != expected {
		t.Errorf("expected value: %v, got value: %v", expected, test)
	}
}

func TestHasSingleCycle4(t *testing.T) {
	expected := true
	array := []int{10, 11, -6, -23, -2, 3, 88, 908, -26}
	test := HasSingleCycle(array)

	if test != expected {
		t.Errorf("expected value: %v, got value: %v", expected, test)
	}
}
