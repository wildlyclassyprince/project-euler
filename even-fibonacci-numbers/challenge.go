package main

import "fmt"

func main() {
	sum := Sum(generateFibonacciNumbers(35))
	fmt.Printf("The sum of even-valued Fibonacci numbers less than four million is %d", sum)
}

func generateFibonacciNumbers(nthTerm int) []int {
	var first, second = 1, 1
	var seq []int
	for i := 0; i < nthTerm; i++ {
		temp := first
		first = second
		second = temp + first
		if first <= 4000000 && first%2 == 0 {
			seq = append(seq, first)
		}
	}
	return seq
}

// Sum returns the sum of the multiples
func Sum(seq []int) int {
	var result int

	for _, v := range seq {
		result += v
	}
	return result
}
