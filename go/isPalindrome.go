package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	length := len(s)

	l := 0
	r := length - 1
	for {
		if l >= r {
			break
		}

		if !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])){
			r--
			continue
		}
		
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

/*
func main() {
	s := "0P"
	ans := isPalindrome(s)
	fmt.Println(ans)
	fmt.Println(unicode.ToLower(rune(s[0])))
	fmt.Println(unicode.ToLower(rune(s[1])))
}
*/