package main

func isPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	for {
		if fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			break
		}
	}

	slow = slow.Next
	slow = reverseList(slow)

	for {
		if slow == nil {
			break
		}
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prev := head
	cur := head.Next

	for {
		if cur == nil {
			break
		}

		t := cur.Next
		cur.Next = prev
		prev = cur
		cur = t
	}
	head.Next = nil
	return prev
}