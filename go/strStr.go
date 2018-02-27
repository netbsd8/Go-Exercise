package main

func strStr(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)

	if ln == 0 {
		return 0
	}

	if lh < ln {
		return -1
	}

	i := 0
	for {
		if i > lh-ln {
			break
		}
		cur := i
		j := 0
		for {
			if j == ln {
				return cur
			}

			if haystack[i] != needle[j] {
				break
			}
			i++
			j++
		}

		i = cur + 1
	}

	return -1
}
