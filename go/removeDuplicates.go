package main

func removeDuplicates(nums []int) int {
	length := len(nums)

	if length == 0 || length == 1 {
		return length
	}

	prev := 0
	cur := 1
	for {
		if cur == length {
			break
		}
		if nums[prev] == nums[cur] {
			cur++
			continue
		}

		prev++
		nums[prev] = nums[cur]
		cur++
	}

	return prev + 1
}
