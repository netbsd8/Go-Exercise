package main

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if isBalanced(root.Left) && isBalanced(root.Right) {
		diff := getTreeHeight(root.Left) - getTreeHeight(root.Right)
		if diff >= 0 && diff <= 1 {
			return true
		}
		if diff <=0 && diff >= -1 {
			return true
		}
		return false
	}
	return false
}

func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(getTreeHeight(root.Left)), float64(getTreeHeight(root.Right))))
}