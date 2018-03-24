package main

import (
	//"fmt"
)

func rotate3(nums []int, k int) {
	length := len(nums)
	k = k%length

	reverseSlice(nums[0:length-k])
	reverseSlice(nums[length-k:])
	reverseSlice(nums)
}

func rotate2(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}

	count := k%length
	if count == 0 {
		return
	}

	t := make([]int,length-count)
	copy(t, nums[0:length-count])
	for i:=0; i<count; i++ {
		nums[i] = nums[length-count+i]
	}

	for i:=0; i<length-count; i++ {
		nums[count+i] = t[i]
	}
}

/*
func main() {
	ans := []int{1, 2}
	rotate2(ans, 1)
	fmt.Println(ans)
}
*/

func rotate1(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}

	count := k%length
	for {
		if count == 0 {
			break
		}

		rotateOne(nums)
		count--

	}
}

func rotateOne(nums []int) {
	length := len(nums)
	t := nums[length-1]
	for i:=length-2; i>=0; i-- {
		nums[i+1] = nums[i]
	}
	nums[0] = t
}

func reverseSlice(nums []int) {
	for l,r :=0,len(nums); l<r; l,r=l+1,r-1 {
		nums[l],nums[r] = nums[r], nums[l]
	}
}