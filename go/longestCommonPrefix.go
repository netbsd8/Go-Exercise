package main

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	if length == 1 {
		return strs[0]
	}

	var ret string
	ret = strs[0]

	for i := 1; i < length; i++ {
		ret = getLongestCommonPrefix(ret, strs[i])
	}

	return ret
}

func getLongestCommonPrefix(str1 string, str2 string) string {
	var ret string
	var minLen int
	len1 := len(str1)
	len2 := len(str2)

	if len1 > len2 {
		minLen = len2
	} else {
		minLen = len1
	}

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			break
		}

		ret = ret + string(str1[i])
	}

	return ret
}

func longestCommonPrefix2(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	if length == 1 {
		return strs[0]
	}

	var ret string
	firstStrLen := len(strs[0])

	for i := 0; i < firstStrLen; i++ {
		curByte := strs[0][i]
		for j := 1; j < length; j++ {
			if i == len(strs[j]) {
				return ret
			}
			if curByte != strs[j][i] {
				return ret
			}
		}
		ret = ret + string(curByte)
	}
	return ret
}
