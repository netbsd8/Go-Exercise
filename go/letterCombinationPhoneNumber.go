package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	m := make(map[byte][]byte)
	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}
	m['0'] = []byte{' '}
	m['*'] = []byte{'+'}

	ans := []string{}
	if len(digits) == 0 {
		return ans
	}

	letterComHelper(&ans, m, digits, "", 0)
	return ans
}

func letterComHelper(ans *[]string, dict map[byte][]byte, digits string, oneAns string, count int) {
	if count == len(digits) {
		*ans = append(*ans, oneAns)
		fmt.Println("ans= ", ans)
		return
	}

	val := dict[digits[count]]
	for j:=0; j<len(val); j++ {
		oneAns = oneAns + string(val[j])
		fmt.Println(oneAns)
		letterComHelper(ans, dict, digits, oneAns, count+1) 
		oneAns = oneAns[:len(oneAns)-1]
	}
}

/*
func main() {
	fmt.Println(letterCombinations("23"))
}
*/