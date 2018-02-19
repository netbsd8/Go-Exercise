package main

import (
	"math"
)

func isUnique(s string, begin int, end int) bool {
	m := make(map[byte]int)
	for i := begin; i <= end; i++ {
		_, ok := m[s[i]]
		if ok {
			return false
		}

		m[s[i]] = 1
	}
	return true
}

// wrong answer
func lengthOfLongestSubString1(s string) int {
	length := len(s)
	var usedArray [256]int
	ret := 0
	for i := 0; i < length; i++ {
		count := 0
		for j := 0; j <= i; j++ {
			count++
			if usedArray[s[j]] == 0 {
				ret = int(math.Max(float64(ret), float64(count)))
				usedArray[s[j]] = 1
			} else {
				usedArray[s[j]] = 0
				count = 0
			}
		}
	}
	return ret
}

// O(n^3)
func lengthOfLongestSubString2(s string) int {
	length := len(s)
	ret := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if !isUnique(s, i, j) {
				continue
			}

			ret = int(math.Max(float64(ret), float64(j-i+1)))
		}
	}
	return ret
}

// O(n)
func lengthOfLongestSubString3(s string) int {
	length := len(s)
	m := make(map[byte]int)
	ret := 0
	j := 0
	for i := 0; i < length; i++ {
		for {
			if j >= length {
				break
			}

			value := m[s[j]]
			if value == 1 {
				break
			}

			ret = int(math.Max(float64(ret), float64(j-i+1)))
			m[s[j]] = 1
			j++
		}
		m[s[i]] = 0
	}
	return ret
}
