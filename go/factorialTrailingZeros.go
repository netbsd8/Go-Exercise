package main

func trailingZeroes2(n int) int {
	ans := 0
	for {
		if n < 5 {
			break
		}
		ans = ans + n/5
		n = n/5
	}

	return ans
}

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	return n/5 + trailingZeroes(n/5)
}