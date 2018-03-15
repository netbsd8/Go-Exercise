package main

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	var reverseValue int
	originalValue := x

	for {
		if x == 0 {
			break
		}
		remain := x % 10
		x = x / 10
		reverseValue = reverseValue*10 + remain
	}

	if reverseValue == originalValue {
		return true
	}
	return false
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	originalValue := x
	div := 1

	for {
		if x < 10 {
			break
		}
		x = x / 10
		div = div * 10
	}

	x = originalValue

	for {
		if x <= 0 {
			break
		}

		leftValue := x / div
		rightValue := x % 10
		if leftValue != rightValue {
			return false
		}

		x = (x % div) / 10
		div = div / 100
		if div == 0 {
			break
		}
	}
	return true
}
