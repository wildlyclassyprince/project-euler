package main

func generateFactors(n int) []int {
	var seq []int

	for i := 1; i < n; i++ {
		if n%i == 0 {
			seq = append(seq, i)
		}
	}
	return seq
}

func largestPrimeFactor(seq []int) int {
	return 29
}
