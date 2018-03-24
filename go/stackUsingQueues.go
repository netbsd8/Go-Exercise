package main

import (
	//"fmt"
)

type MyStack struct {
	q		chan int
	top		int
}


/** Initialize your data structure here. */
func StackConstructor() MyStack {
	s := MyStack{}
	s.q = make(chan int, 10)
	s.top = 0
	return s
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.q <- x
	this.top = x
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	q2 := make(chan int, 10)
	length := len(this.q)
	for i:=0; i<length-1; i++ {
		x := <- this.q
		q2 <- x
		if i == length-2 {
			this.top = x
		}
	}

	length = len(q2)
	for i:=0; i<length; i++ {
		x := <- q2
		this.q <- x
	}

	ans := <- this.q
	return ans
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
   return len(this.q) == 0 
}

/*
func main() {
	s := StackConstructor();
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}
*/

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */