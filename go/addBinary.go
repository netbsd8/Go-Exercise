package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)

	if la == 0 && lb == 0 {
		return ""
	}
	if la == 0 {
		return b
	}
	if lb == 0 {
		return a
	}

	ret := ""

	la = la - 1
	lb = lb - 1
	var carry int
	for {
		if la < 0 || lb < 0 {
			break
		}
		cur := ((int(a[la]) - '0') + (int(b[lb]) - '0') + carry) % 2
		carry = ((int(a[la]) - '0') + (int(b[lb]) - '0') + carry) / 2
		ret = strconv.Itoa(cur) + ret
		la--
		lb--
	}

	for {
		if la < 0 {
			break
		}
		cur := (int(a[la]) - '0' + carry) % 2
		carry = (int(a[la]) - '0' + carry) / 2
		ret = strconv.Itoa(cur) + ret
		la--
	}

	for {
		if lb < 0 {
			break
		}
		cur := (int(b[lb]) - '0' + carry) % 2
		carry = (int(b[lb]) - '0' + carry) / 2
		ret = strconv.Itoa(cur) + ret
		lb--
	}

	if carry == 1 {
		ret = strconv.Itoa(carry) + ret
	}

	return ret
}
