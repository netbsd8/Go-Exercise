package main

func majorityElement(nums []int) int {
	ans := nums[0]
	count := 1
	length := len(nums)

	for i:=1; i<length; i++ {
		if ans != nums[i] {
			count--
			if count == 0 {
				ans = nums[i]
				count = 1
			}
		} else {
			count++
			if count > length/2 {
				break
			}
		}
	}

	return ans
}