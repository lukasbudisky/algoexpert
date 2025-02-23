package main

import "testing"

func TestGenerateDocument1(t *testing.T) {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"

	expected := true
	test := GenerateDocument(characters, document)
	if test != expected {
		t.Errorf("expected output: %v, got output: %v", expected, test)
	}
}

func TestGenerateDocument2(t *testing.T) {
	characters := "Bste!hetsi ogEAxpert"
	document := "AlgoExpert is the Best!"

	expected := false
	test := GenerateDocument(characters, document)
	if test != expected {
		t.Errorf("expected output: %v, got output: %v", expected, test)
	}
}
