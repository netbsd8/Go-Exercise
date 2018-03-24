package main

func moveZeroesNoOrder(nums []int) {
	length := len(nums)
	l := 0
	r := length-1

	for {
		if l >= r {
			break
		}

		if nums[l] == 0 {
			t := nums[r]
			nums[r] = nums[l]
			nums[l] = t
			r--
		} else {
			l++
		}
	}
}

func moveZeroesOrder(nums []int) {
	length := len(nums)
	prev := 0
	cur := 0 

	for {
		if cur >= length {
			break
		}

		if nums[cur] == 0 {
			cur++
			continue
		}

		t := nums[prev]
		nums[prev] = nums[cur]
		nums[cur] = t
		prev++
		cur++
	}
}