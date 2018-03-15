package main

import (
	"math"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := make(chan *TreeNode)
	q <- root
	count := 0
	var cur *TreeNode
	for {
		if len(q) == 0 {
			break
		}

		count++

		for i:=0; i<len(q); i++ {
			cur =<- q
			if cur.Left != nil {
				q <- cur.Left
			}
			if cur.Right != nil{
				q <- cur.Right
			}
		}
	}
	return count
}