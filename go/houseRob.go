package main

import (
	"math"
	//"fmt"
)

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	withNB := 0
	withoutNB := 0

	for i:=0; i<length; i++ {
		t := withNB
		withNB = int(math.Max(float64(withoutNB), float64(withNB)))
		withoutNB = int(math.Max(float64(t + nums[i]), float64(withoutNB)))
	}

	if withNB > withoutNB {
		return withNB
	}
	return withoutNB
}

func robDp(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i:=2; i<length; i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+nums[i]), float64(dp[i-1])))
	}
	return dp[length-1]
}

/*
func main() {
	a := []int{5, 10, 2, 4,10}
	fmt.Println(robDp(a))
}
*/