package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	ans := [][]int{}
	if n == 0 || k == 0 {
		 return ans
	}

	combineHelper(&ans, &[]int{}, n, k, 1)

	return ans

}

func combineHelper(ans *[][]int, oneAns *[]int, n int, k int, start int) {
	if len(*oneAns) == k {
		fmt.Println("before appending to ans ", *oneAns)
		*ans = append(*ans, *oneAns)
		fmt.Println("ans ", *ans)
		return
	}

	for i:=start; i<=n; i++ {
		fmt.Println("i vs n vs count ", i, n, start)
		*oneAns = append(*oneAns, i)
		fmt.Println(*oneAns)
		combineHelper(ans, oneAns, n, k, i+1)
		*oneAns = (*oneAns)[:len(*oneAns)-1] 
	}
}

/*
func main() {
	fmt.Println("Result ", combine(2,2))
}
*/