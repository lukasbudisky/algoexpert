package main

import "testing"

func TestLevenshteinDistance1(t *testing.T) {
	str1 := "abc"
	str2 := "yabd"
	result := 2

	got := LevenshteinDistance(str1, str2)
	if result != got {
		t.Errorf("expected result: %d, got result: %d", result, got)
	}
}

func TestLevenshteinDistance2(t *testing.T) {
	str1 := "algoexpert"
	str2 := "algozexpert"
	result := 1

	got := LevenshteinDistance(str1, str2)
	if result != got {
		t.Errorf("expected result: %d, got result: %d", result, got)
	}
}

func TestLevenshteinDistance3(t *testing.T) {
	str1 := ""
	str2 := "abc"
	result := 3

	got := LevenshteinDistance(str1, str2)
	if result != got {
		t.Errorf("expected result: %d, got result: %d", result, got)
	}
}

func TestLevenshteinDistance4(t *testing.T) {
	str1 := "biting"
	str2 := "mitten"
	result := 4

	got := LevenshteinDistance(str1, str2)
	if result != got {
		t.Errorf("expected result: %d, got result: %d", result, got)
	}
}
