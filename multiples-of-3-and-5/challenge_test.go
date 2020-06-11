package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of all multiples 3 or 5 below 10", func(t *testing.T) {
		got := Sum(3, 5, 10)
		want := 23
		assertSum(t, got, want)
	})
}

func assertSum(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
