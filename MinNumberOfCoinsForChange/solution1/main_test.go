package main

import "testing"

func TestMinNumberOfCoinsForChange_1(t *testing.T) {
	n := 7
	denoms := []int{2, 4}
	expect := -1

	result := MinNumberOfCoinsForChange(n, denoms)
	if result != expect {
		t.Errorf("expect value: %d, got value: %d", expect, result)
	}

}

func TestMinNumberOfCoinsForChange_2(t *testing.T) {
	n := 9
	denoms := []int{3, 5}
	expect := 3

	result := MinNumberOfCoinsForChange(n, denoms)
	if result != expect {
		t.Errorf("expect value: %d, got value: %d", expect, result)
	}
}

func TestMinNumberOfCoinsForChange_3(t *testing.T) {
	n := 6
	denoms := []int{1, 2, 4}
	expect := 2

	result := MinNumberOfCoinsForChange(n, denoms)
	if result != expect {
		t.Errorf("expect value: %d, got value: %d", expect, result)
	}
}

func TestMinNumberOfCoinsForChange_4(t *testing.T) {
	n := 135
	denoms := []int{39, 45, 130, 40, 4, 1, 60, 75}
	expect := 2

	result := MinNumberOfCoinsForChange(n, denoms)
	if result != expect {
		t.Errorf("expect value: %d, got value: %d", expect, result)
	}
}

func TestMinNumberOfCoinsForChange_5(t *testing.T) {
	n := 3
	denoms := []int{2, 1}
	expect := 2

	result := MinNumberOfCoinsForChange(n, denoms)
	if result != expect {
		t.Errorf("expect value: %d, got value: %d", expect, result)
	}
}

func TestMinNumberOfCoinsForChange_6(t *testing.T) {
	n := 7
	denoms := []int{1, 5, 10}
	expect := 3

	result := MinNumberOfCoinsForChange(n, denoms)
	if result != expect {
		t.Errorf("expect value: %d, got value: %d", expect, result)
	}
}
