package main

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if target > nums[length-1] {
		return length
	}

	if target < nums[0] {
		return 0
	}

	l := 0
	r := length - 1

	for {
		if l+1 >= r {
			break
		}

		m := l + (r-l)/2
		if target == nums[m] {
			return m
		}

		if target > nums[m] {
			l = m
		}

		if target < nums[m] {
			r = m
		}
	}

	if target <= nums[l] {
		return l
	}
	if target <= nums[r] {
		return r
	}
	return r + 1
}
