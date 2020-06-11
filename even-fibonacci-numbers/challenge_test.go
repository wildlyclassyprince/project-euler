package main

import "testing"

func TestGeneratedFibonacciNumbers(t *testing.T) {
	got := generateFibonacciNumbers(10)
	want := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	assertSequence(t, got, want)
}

func TestSum(t *testing.T) {
	t.Run("test sum of the first 10 even-valued Fibonacci numbers", func(t *testing.T) {
		seq := generateFibonacciNumbers(10)
		got := Sum(seq)
		want := 44
		assertSum(t, got, want)
	})
	t.Run("test sum of even-valued Fibonacci numbers less than 4 million", func(t *testing.T) {
		seq := generateFibonacciNumbers(35)
		got := Sum(seq)
		want := 4613732
		assertSum(t, got, want)
	})
}

func assertSequence(t *testing.T, got, want []int) {
	t.Helper()

	if got == nil {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertSum(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
