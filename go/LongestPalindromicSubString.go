package main

// DP: O(n)
func longestPalindrome1(s string) string {
	length := len(s)

	var dp [1000][1000]int
	for m := 0; m < 1000; m++ {
		for n := 0; n < 1000; n++ {
			dp[m][n] = 0
		}
	}
	maxLen := 0
	l := 0
	r := 0

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if s[i] != s[j] {
				continue
			}

			if j+1 == i || dp[j+1][i-1] == 1 {
				dp[j][i] = 1

				if maxLen >= i-j+1 {
					continue
				}
				maxLen = i - j + 1
				l = j
				r = i
			}
		}
		dp[i][i] = 1
	}
	return s[l : r+1]
}

// O(n^2)
func longestPalindrome2(s string) string {
	maxStr := s[0:1]
	length := len(s)

	for i := 0; i < length-1; i++ {
		curStr := findPalindrome(s, i, i)
		if len(maxStr) < len(curStr) {
			maxStr = curStr
		}

		if s[i] != s[i+1] {
			continue
		}

		curStr = findPalindrome(s, i, i+1)
		if len(maxStr) < len(curStr) {
			maxStr = curStr
		}
	}
	return maxStr
}

func findPalindrome(s string, left int, right int) string {
	for {
		left--
		right++
		if left < 0 || right >= len(s) {
			break
		}
		if s[left] != s[right] {
			break
		}
	}
	return s[left+1 : right]
}
