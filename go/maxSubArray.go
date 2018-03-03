package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	ret := nums[0]
	curMax := nums[0]
	length := len(nums)

	for i := 1; i < length; i++ {
		curMax = int(math.Max(float64(curMax+nums[i]), float64(nums[i])))

		ret = int(math.Max(float64(ret), float64(curMax)))
	}

	return ret
}
