package main

import (
	"fmt"
	"math"
)

func missingNumber(nums []int) int {
	length := len(nums)
	min := math.MaxInt32
	max := math.MinInt32

	sum := 0
	for i:=0; i<length; i++ {
		min = int(math.Min(float64(min), float64(nums[i])))
		max = int(math.Max(float64(max), float64(nums[i])))
		sum = sum + nums[i]
	} 

	e := 0
	fmt.Println("min ", min)
	fmt.Println("max ", max)
	for i:=min; i<=max; i++ {
		e = e + i
		fmt.Println("i ", i)
	}
	fmt.Println("e ", e)
	fmt.Println("sum ", sum)

	return e-sum
}


func missingNumber2(nums []int) int {
	length := len(nums)
	min := math.MaxInt32
	max := math.MinInt32

	sum := 0
	for i:=0; i<length; i++ {
		min = int(math.Min(float64(min), float64(nums[i])))
		max = int(math.Max(float64(max), float64(nums[i])))
		sum = sum ^ nums[i]
	}
	
	newSum := min 
	for i:=min+1; i<=max; i++ {
		newSum = newSum ^ i
	}
	return newSum^sum
}

/*
func main() {
	fmt.Println(missingNumber2([]int{3, 0, 1}))
}
*/