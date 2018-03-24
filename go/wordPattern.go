package main

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	la := len(pattern)
	strs := strings.Split(str, " ")
	lb := len(strs)
	
	if la != lb {
		return false
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	for i:=0; i<la; i++ {
		val1, ok1 := m1[pattern[i]]
		val2, ok2 := m2[strs[i]]
		if !ok1 && !ok2 {
			m1[pattern[i]] = strs[i]
			m2[strs[i]] = pattern[i]
		} else if ok1 && ok2 {
			if val1 != strs[i] || val2 != pattern[i] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}