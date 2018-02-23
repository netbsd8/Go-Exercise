package main

func removeNthFormEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head

	pprev := dummyNode
	prev := head
	cur := head
	for i := 0; i < n; i++ {
		cur = cur.Next
	}

	for {
		if cur.Next == nil {
			break
		}
		pprev = pprev.Next
		prev = prev.Next
		cur = cur.Next
	}

	pprev.Next = prev.Next

	return dummyNode.Next
}
