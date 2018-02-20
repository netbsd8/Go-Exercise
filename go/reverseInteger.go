package main

const (
	MAX int = 1<<31 - 1
	MIN     = -MAX - 1
)

func reverse(x int) int {
	ret := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	for {
		y := x % 10
		ret = ret + y

		if ret > MAX {
			return 0
		}

		if x < 10 {
			break
		}

		x = x / 10
		ret = ret * 10
	}

	return ret * sign
}
