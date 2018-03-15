package main

import (
	"math"
)

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	ans := math.MinInt32
	curMin := prices[0]
	for i:=0; i<length; i++ {
		curMax := prices[i] - curMin
		curMin = int(math.Min(float64(curMin), float64(prices[i])))

		ans = int(math.Max(float64(ans), float64(curMax)))
	}

	return ans
}