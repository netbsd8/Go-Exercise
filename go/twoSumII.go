package main

func twoSumII(numbers []int, target int) []int {
	length := len(numbers)

	l := 0
	r := length-1
	ans := []int{}

	for {
		if l >= r {
			break
		}

		sum := numbers[l] + numbers[r]

		if sum == target {
			ans = append(ans, l+1)
			ans = append(ans, r+1)
			return ans
		}

		if sum > target {
			r--
		} else {
			l++
		}
	}

	return ans 
}