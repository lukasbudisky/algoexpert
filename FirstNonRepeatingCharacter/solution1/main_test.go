package main

import "testing"

func TestFirstNonRepeatingCharacter1(t *testing.T) {
	s := "faadabcbbebdf"
	expected := 6

	test := FirstNonRepeatingCharacter(s)
	if test != expected {
		t.Errorf("expected output: %d, got output: %d", expected, test)
	}
}

func TestFirstNonRepeatingCharacter2(t *testing.T) {
	s := "abcdcaf"
	expected := 1

	test := FirstNonRepeatingCharacter(s)
	if test != expected {
		t.Errorf("expected output: %d, got output: %d", expected, test)
	}
}
