package main

func lengthOfLastWord(s string) int {
	str := s
	i := len(str) - 1
	for ; i >= 0; i-- {
		if string(str[i]) != " " {
			break
		}
	}
	str = str[0 : i+1]

	length := len(str)
	r := length - 1
	for {
		if r < 0 {
			break
		}

		if string(s[r]) == " " {
			break
		}
		r--
	}

	return (length - 1) - r
}
