package main

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	length := len(nums)
	m := make(map[int]bool)
	for i:=0; i<length; i++ {
		m[i] = false
	}

	for i:=0; i<length; i++ {
		ok := m[nums[i]]
		if ok {
			return true
		}
		m[nums[i]] = true

	}
	return false
}

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)

	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}