package main

import (
	//"fmt"
)

func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}

	cnt := 0
	for i:=0; i<31; i++ {
		if n&1 == 1 {
			cnt++
		}
		n = n >> 1
	}
	return cnt == 1
}

func isPowerOfTwo2(n int) bool {
	return n>0 && (n&(n-1)) == 0
}

/*
func main() {
	fmt.Println(isPowerOfTwo2(2))
}
*/