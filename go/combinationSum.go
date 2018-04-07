package main

import (
	"sort"
)

func combinationSum1(candidates []int, target int) [][]int {
	ans := [][]int{}

	sort.Ints(candidates)  // nedd it? yes, otherwise it could not avoid repeating

	combinationSumHelper1(ans, []int{}, candidates, target, 0, 0)

	return ans
}

func combinationSumHelper1(ans [][]int, oneAns []int, candidates []int, target int, sum int, start int) {
	if sum == target {
		ans = append(ans, oneAns)
		return
	}
	if sum > target {
		return
	}

	for i:=start; i<len(candidates); i++ {
		c := candidates[i]
		newAns := make([]int, len(oneAns))
		copy(newAns, oneAns)
		newAns = append(newAns, c) 
		combinationSumHelper1(ans, newAns, candidates, target, sum+c, i)
		newAns = newAns[:len(newAns)-1]
	}
}