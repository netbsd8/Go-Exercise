package main

// ListNode: comment to export
type ListNode struct {
	Val  int
	Next *ListNode
}

func add2nums1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := ListNode{}
	head := &ListNode{1, nil}
	dummy.Next = head
	carry := 0
	for {
		if l1.Next != nil && l2.Next != nil {
			sum := l1.Val + l2.Val + carry
			head.Val = sum % 10
			if sum >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			newNode := ListNode{1, nil}
			head.Next = &newNode
			head = head.Next
			l1 = l1.Next
			l2 = l2.Next
		} else {
			break
		}
	}

	for {
		if l1.Next != nil {
			sum := l1.Val + carry
			head.Val = sum % 10
			if sum >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			newNode := ListNode{1, nil}
			head.Next = &newNode
			head = head.Next
			l1 = l1.Next
		} else {
			break
		}
	}

	for {
		if l2.Next != nil {
			sum := l2.Val + carry
			head.Val = sum % 10
			if sum >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			newNode := ListNode{1, nil}
			head.Next = &newNode
			head = head.Next
			l2 = l2.Next
		} else {
			break
		}
	}

	if carry == 1 {
		newNode := ListNode{1, nil}
		head.Next = &newNode
	}

	return dummy.Next
}

func add2nums2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for {
		if l1 != nil || l2 != nil {
			var n1 int
			var n2 int

			if l1 == nil {
				n1 = 0
			} else {
				n1 = l1.Val
			}

			if l2 == nil {
				n2 = 0
			} else {
				n2 = l2.Val
			}

			sum := n1 + n2 + carry
			carry = sum / 10

			newNode := ListNode{sum % 10, nil}
			cur.Next = &newNode
			cur = cur.Next
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		} else {
			break
		}
	}

	if carry == 1 {
		newNode := ListNode{1, nil}
		cur.Next = &newNode
	}

	return dummy.Next
}
