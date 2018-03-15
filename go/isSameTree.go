package main

import (
	"github.com/golang-collections/collections/stack"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	s := stack.New()

	s.Push(p)

	// pre-order traversal
	for {
		if s.Len() == 0 && p == nil {
			break
		}
	
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			p = s.Pop().(*TreeNode)
			p = p.Right
		}
	}

	/*
	// another pre-order traversal
	for {
		if s.Len() == 0 {
			break
		}

		// access p
		p = s.Pop().(*TreeNode)

		if p.Right != nil {
			s.Push(p.Right)
		}
		if p.Left != nil {
			s.Push(p.Left)
		}
	}
	*/

	return true
}
