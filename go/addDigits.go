package main

func addDigits(num int) int {
	for {
		if num < 10 {
			break
		}

		sum := 0
		for {
			if num == 0 {
				break
			}
			sum = sum + num%10
			num = num/10
		}
		num = sum
	}

	return num
}