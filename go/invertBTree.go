package main

import (
	"github.com/golang-collections/collections/queue"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	t := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(t)
	return root	
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := queue.New()
	q.Enqueue(root)

	for {
		if q.Len() == 0 {
			break	
		}
		
		cur := q.Dequeue().(*TreeNode)
		t := cur.Left
		cur.Right = cur.Left
		cur.Left = t

		if cur.Left != nil {
			q.Enqueue(cur.Left)
		}
		if cur.Right != nil {
			q.Enqueue(cur.Right)
		}
	}

	return root
}