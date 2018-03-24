package main

func canWinNim(n int) bool {
	if n <= 3 {
		return true
	}

	if n == 4 {
		 return false
	}

	return canWinNim(n-4)
}