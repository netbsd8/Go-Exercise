package main

import (
	"math"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	return 1 + int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right))))
}