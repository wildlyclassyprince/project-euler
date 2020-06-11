package main

import "fmt"

func main() {
	m1, m2 := 3, 5
	limit := 1000
	fmt.Printf("Sum of all multiples of %d and %d below %d is %d.", m1, m2, limit, Sum(m1, m2, limit))
}

// Sum returns the sum of the multiples
func Sum(m1, m2, limit int) int {
	var i, result int
	var list []int

	for i = 0; i < limit; i++ {
		if (i%m1 == 0) || (i%m2 == 0) {
			list = append(list, i)
		}
	}
	for _, v := range list {
		result += v
	}
	return result
}
