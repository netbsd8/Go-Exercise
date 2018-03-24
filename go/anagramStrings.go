package main

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}

	ns := sortStringByChar(s)
	nt := sortStringByChar(t)

	for i:=0; i<ls; i++ {
		if ns[i] != nt[i] {
			return false
		}
	}
	return true
}

func sortStringByChar(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str,"")
}


func isAnagram2(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}

	m := make(map[byte]int)
	for i:=0; i<ls; i++ {
		_, ok := m[s[i]]
		if !ok {
			m[s[i]] = 1
		} else {
			m[s[i]]++
		}
	}

	for i:=0; i<lt; i++ {
		val, ok := m[t[i]]
		if val < 0 || !ok {
			return false
		}
		m[t[i]]--;
	}

	for _, v := range m {
		if v != 0 {
			return false	
		} 
	}
	return true
}