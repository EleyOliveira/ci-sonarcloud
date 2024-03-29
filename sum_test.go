package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("O resultado deve ser 5")
	}
}

func TestSub(t *testing.T) {
	result := sub(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestTimes(t *testing.T) {
	result := times(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestDiv(t *testing.T) {
	result := div(10, 2)
	if result != 5.0 {
		t.Errorf("Expected 5.0, but got %.2f", result)
	}

	// Test division by zero
	result = div(5, 0)
	if result != 0.0 {
		t.Errorf("Expected 0.0, but got %.2f", result)
	}
}
