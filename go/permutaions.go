package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}

	m := make(map[int]bool)
	permuteHelper(&ans, nums, []int{}, 0, &m)
	return ans
}

func permuteHelper(ans *[][]int, nums []int, oneAns []int, count int, dict *map[int]bool) {
	if count == len(nums) {
		fmt.Println("oneAns", oneAns)
		fmt.Println("ans ", *ans)
		*ans = append(*ans, oneAns)
		fmt.Println("ans after insert ", *ans)
		return
	}

	for i:=0; i<len(nums); i++ {
		if !(*dict)[nums[i]] {
			(*dict)[nums[i]] = true
			oneAns = append(oneAns, nums[i])
			permuteHelper(ans, nums, oneAns, count+1, dict)
			oneAns = oneAns[:len(oneAns)-1]
			(*dict)[nums[i]] = false
		}
	}
}

/*
func main() {
	fmt.Println(permute([]int{5, 4, 6}))
}
*/