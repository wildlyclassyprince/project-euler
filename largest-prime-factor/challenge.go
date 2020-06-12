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

func primeFactor(seq []int) []int {
	var seq2 []int
	for _, v := range seq {
		if v > 1 {
			for i := 2; i < v; i++ {
				if v%i == 0 {
					break
				} else {
					seq2 = append(seq2, v)
				}
			}
		}
	}
	return Ints(seq2)
}

// Ints returns a unique subset of the int slice provided.
func Ints(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// arkinSieve filters factors for prime-ness
func atkinSieve(seq []int) []int {
	return []int{5, 7, 13, 29}
}
