package main

import "math"

func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}

	np := make([]bool, n)
	for i:=0; i<n; i++ {
		np[i] = true
	}
	np[0] = false
	np[1] = false

	for i:=2; i<=int(math.Sqrt(float64(n))); i++ {
		if np[i] == false {
			continue
		}

		for j:=i*i; j<n; j=j+i {
			np[j] = false
		}	
	}

	ans := 0
	for i:=0; i<n; i++ {
		if np[i] == true {
			ans++
		}
	}

	return ans
}