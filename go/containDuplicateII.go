package main

func containsDuplicateII(nums []int, k int) bool {
	length := len(nums)
	m := make(map[int]int)

	for i:=0; i<length; i++ {
		v,ok := m[nums[i]]
		if !ok {
			m[nums[i]] = i
			continue
		}

		if i-v <= k {
			return true
		} else {
			m[nums[i]] = i
		}
	}

	return false
}