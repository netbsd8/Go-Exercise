package main 

import (
	//"fmt"
)

func getRow(rowIndex int) []int {
	ans := []int{1}

	if rowIndex == 0 {
		return ans
	} 

	for i:=1; i<=rowIndex; i++ {
		ans = append(ans, 1)
		ans[0] = 1
		ans[i] = ans[1]

		for j:=1; j<i; j++ {
			ans[j] = ans[0] + ans[i]
			ans[0] = ans[i]
			ans[i] = ans[j+1]
		}

		ans[0] = 1
		ans[i] = 1
	}

	return ans
}

/*
func main() {
	ans := getRow(1)
	fmt.Println(ans)
}
*/