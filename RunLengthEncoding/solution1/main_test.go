package main

import "testing"

func TestRunLengthEncoding1(t *testing.T) {
	s := "AAAAAAAAAAAAABBCCCCDD"
	expected := "9A4A2B4C2D"

	test := RunLengthEncoding(s)
	if test != expected {
		t.Errorf("expected output: %s, got output: %s", expected, test)
	}
}

func TestRunLengthEncoding2(t *testing.T) {
	s := "........______=========AAAA   AAABBBB   BBB"
	expected := "8.6_9=4A3 3A4B3 3B"

	test := RunLengthEncoding(s)
	if test != expected {
		t.Errorf("expected output: %s, got output: %s", expected, test)
	}
}
