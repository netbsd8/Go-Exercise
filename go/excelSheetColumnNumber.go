package main

import (
	//"fmt"
)

func titleToNumbers(s string) int {
	ans := 0
	rs := []rune(s)
	for i:=0; i<len(rs); i++ {
		ans = ans * 26
		ans = ans + int(rs[i] - 'A' + 1)
	}

	return ans
}

/*
func main() {
	fmt.Println(titleToNumbers("A"))
	fmt.Println(titleToNumbers("B"))
	fmt.Println(titleToNumbers("Z"))
	fmt.Println(titleToNumbers("AA"))
	fmt.Println(titleToNumbers("AB"))
	fmt.Println(titleToNumbers("AZ"))
	fmt.Println(titleToNumbers("BA"))
	fmt.Println(titleToNumbers("BB"))
}
*/