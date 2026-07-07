package main

import "testing"

func TestDoMath(t *testing.T) {
	want := 58

	got := doMath(42, 16, func(a, b int) int {
		return a + b
	})

	if got != want {
		t.Errorf("got: %d; want %d", got, want)
	}
}

func TestAdd(t *testing.T) {
	want := 10

	got := add(5, 5)

	if got != want {
		t.Errorf("got: %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	want := 5

	got := subtract(10, 5)

	if got != want {
		t.Errorf("got: %d; want %d", got, want)
	}
}
