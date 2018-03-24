package main

import (
	//"fmt"
	"strings"
)

func convertToTitle(n int) string {
	ans := ""

	cur := n
	ret := 0
	for {
		ret = (cur-1)%26
		ans = stringInsert(ans, rune('A' + ret)) 
		cur = (cur-ret)/26
		if cur < 1 {
			break;
		}
	}

	/*
	if ret != 0 {
		ans = stringInsert(ans, rune('A' + ret))
	}
	*/

	return ans
}

func stringInsert(s string, c rune) string {
	newStrs := []string{string(c), s}
	return strings.Join(newStrs, "")
}

/*
func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(2))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(26*26))
}
*/