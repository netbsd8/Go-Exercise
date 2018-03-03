package main

import (
	"strconv"
)

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}

	res := "1"
	i := 1
	for {
		if i >= n {
			break
		}
		cur := res
		res = ""
		prev := 0
		j := 0
		for {
			if j >= len(cur) {
				break
			}
			for {
				if j >= len(cur) {
					break
				}
				if cur[j] != cur[prev] {
					break
				}
				j++
			}
			count := j - prev
			res = res + strconv.Itoa(count) + string(cur[prev])
			prev = j
		}
		i++
	}
	return res
}
