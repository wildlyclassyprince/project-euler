package main

import "fmt"

func main() {
	const limit = 600851475143
	var nums []int
	primes := make(chan int)
	go primeSieve(primes)

	p := <-primes
	for p < limit {
		if limit%p == 0 {
			nums = append(nums, p)
		}
		p = <-primes
	}
	fmt.Printf("%v", nums)
	fmt.Println()
	fmt.Printf("Max: %d", findMax(nums))
}

func primeSieve(out chan int) {
	out <- 2
	out <- 3

	primes := make(chan int)
	go primeSieve(primes)

	var p int
	p = <-primes
	p = <-primes

	sieve := make(map[int]int)
	q := p * p
	n := p

	for {
		n += 2
		step, isComposite := sieve[n]
		if isComposite {
			delete(sieve, n)
			m := n + step
			for sieve[m] != 0 {
				m += step
			}
			sieve[m] = step

		} else if n < q {
			out <- n

		} else {
			step = p + p
			m := n + step
			for sieve[m] != 0 {
				m += step
			}
			sieve[m] = step
			p = <-primes
			q = p * p
		}
	}
}

func findMax(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[0] < nums[i] {
			nums[0] = nums[i]
		}
	}
	return nums[0]
}
