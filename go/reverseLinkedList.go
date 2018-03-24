package main

/*
type ListNode struct {
	Val  int
	Next *ListNode
}
*/


func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var prev *ListNode

	for {
		if cur == nil {
			break
		}

		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	head.Next = nil
	return prev
}
