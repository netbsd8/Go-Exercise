package main

import (
	//"fmt"
	stack "github.com/golang-collections/collections/stack"
)

type MyQueue struct {
	s1		stack.Stack
	s2		stack.Stack
	peek	int
}


/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	q := MyQueue{}
	q.s1 = stack.Stack{}
	q.s2 = stack.Stack{}
	return q
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
   if this.s1.Len() == 0 {
	   this.peek = x 
   } 
   this.s1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
   length := this.s1.Len()
   for i:=0; i<length; i++ {
	   this.s2.Push(this.s1.Pop())
   }
   ans := this.s2.Pop().(int)
   
   length = this.s2.Len()
   for i:=0; i<length; i++ {
	   x := this.s2.Pop()
	   this.s1.Push(x)
	   if i == 0 {
		   this.peek = x.(int) 
	   }
   }
   return ans
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.peek
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.s1.Len() == 0
}

/*
func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}
*/

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
