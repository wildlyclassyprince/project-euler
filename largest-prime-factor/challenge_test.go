package main

import (
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	t.Run("test all factors of 13195", func(t *testing.T) {
		got := generateFactors(13195)
		want := []int{1, 5, 7, 13, 29, 35, 65, 91, 145, 203, 377, 455, 1015, 1885, 2639}
		assertSequence(t, got, want)
	})
	t.Run("test prime-ness of factors of 13195", func(t *testing.T) {
		seq := generateFactors(13195)
		got := atkinSieve(seq)
		want := []int{5, 7, 13, 29}
		assertSequence(t, got, want)
	})
}

func assertSequence(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertLargestPrime(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
