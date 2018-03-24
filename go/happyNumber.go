package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)
	m[n] = true
	for {
		ans := 0
		for {
			if n == 0 {
				break
			}
			ans = ans + (n%10)*(n%10)
			n = n/10
		}
		fmt.Println(ans)
		if ans == 1 {
			return true
		}
		n = ans
		if m[ans] {
			break
		} else {
			m[ans] = true
		}
	}
	return false
}

/*
func main() {
	fmt.Println(isHappy(3))
}
*/