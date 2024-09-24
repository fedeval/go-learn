package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(5, 5)
	if sum != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", sum, 10)
	}
}

func TestSubtract(t *testing.T) {
	diff := Subtract(10, 5)
	if diff != 5 {
		t.Errorf("Subtraction was incorrect, got: %d, want: %d", diff, 5)
	}
}

func TestDoOperation(t *testing.T) {
	doSum := DoOperation(5, 5, Add)
	if doSum != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", doSum, 10)
	}

	doSub := DoOperation(5, 5, Subtract)
	if doSub != 0 {
		t.Errorf("Subtraction was incorrect, got: %d, want: %d", doSub, 0)
	}
}
