package main

import (
	"math"
)

type MinStack struct {
	s 		[]int
	m 		[]int
	min 	int
}

func Constructor() MinStack {
	st := MinStack{}
	st.min = math.MaxInt32
	return st
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	if this.min > x {
		this.m = append(this.m, x)
		this.min = x
	} else {
		this.m = append(this.m, this.min)
	}
}

func (this *MinStack) Pop() {
	if len(this.s) == 0 {
		return
	}
	this.s = this.s[:len(this.s)-1]
	this.m = this.m[:len(this.m)-1]
	this.min = this.m[len(this.m)-1]
}

func (this *MinStack) Top() int {
	if len(this.s) == 0 {
		return math.MinInt32
	}
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.m) > 0 {
		return this.min
	}
	return math.MaxInt32
}