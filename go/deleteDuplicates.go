package main

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	prev := head
	for {
		if cur == nil {
			break
		}

		for {
			if cur == nil {
				break
			}
			if cur.Val != prev.Val {
				break
			}
			cur = cur.Next
		}
		prev.Next = cur
		prev = prev.Next
		if cur != nil {
			cur = cur.Next
		}
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	cur := head
	for {
		if cur == nil || cur.Next == nil {
			break
		}

		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
