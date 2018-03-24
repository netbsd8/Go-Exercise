package main

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := ListNode{}
	dummyNode.Next = head
	prev := &dummyNode
	cur := head
	
	for {
		if cur == nil {
			prev.Next = cur
			break
		}

		if cur.Val == val {
			cur = cur.Next
			continue
		}

		prev.Next = cur
		prev = prev.Next
		cur = cur.Next
	}

	return dummyNode.Next
}