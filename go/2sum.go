package main

import (
	"sort"
)

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		another := target - num
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[num] = i
	}
	return nil
}

// this is an incorrect solution, since return needs to be index but not value
// so the sort will mess up the original index
func twoSum2(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}

	sort.Ints(nums)
	l := 0
	r := len(nums) - 1

	for {
		if l == r {
			break
		}

		if nums[l]+nums[r] == target {
			return []int{l, r}
		}

		if nums[l]+nums[r] < target {
			l++
		} else {
			r--
		}
	}

	return []int{}
}
