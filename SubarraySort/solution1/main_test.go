package main

import "testing"

func TestCmpValuesLow(t *testing.T) {
	low := 5
	big := 10
	value := 3

	l, b := cmpValues(low, big, value)
	if l != value || b != big {
		t.Errorf("expected low value: %d, got: %d\nexpected high value: %d, got:%d", value, l, big, b)
	}
}

func TestCmpValuesHigh(t *testing.T) {
	low := 5
	big := 10
	value := 11

	l, b := cmpValues(low, big, value)
	if l != low || b != value {
		t.Errorf("expected low value: %d, got: %d\nexpected high value: %d, got:%d", value, l, big, b)
	}
}

func TestSubarraySort1(t *testing.T) {
	array := []int{4, 8, 7, 12, 11, 9, -1, 3, 9, 16, -15, 51, 7}
	expected := []int{0, 12}

	result := SubarraySort(array)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestSubarraySort2(t *testing.T) {
	array := []int{2, 1}
	expected := []int{0, 1}

	result := SubarraySort(array)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestSubarraySort3(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 18, 21, 22, 7, 14, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 4, 14, 11, 6, 33, 35, 41, 55}
	expected := []int{4, 24}

	result := SubarraySort(array)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestSubarraySort4(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{-1, -1}

	result := SubarraySort(array)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestSubarraySort5(t *testing.T) {
	array := []int{1, 2}
	expected := []int{-1, -1}

	result := SubarraySort(array)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestSubarraySort6(t *testing.T) {
	array := []int{1, 2, 8, 4, 5}
	expected := []int{2, 4}

	result := SubarraySort(array)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestSubarraySort7(t *testing.T) {
	array := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	expected := []int{3, 9}

	result := SubarraySort(array)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
