package main

func singleNumber(nums []int) int {
	length := len(nums)
	ans := nums[0]

	for i:=1; i<length; i++ {
		ans = ans ^ nums[i]
	}

	return ans
}