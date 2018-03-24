package main

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	t := "" 
	s := []string{}
	if root != nil {
		btpHelper(root, &s, t)
	}
	return s
}

func btpHelper(root *TreeNode, s *[]string, t string) {
	t = t + strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*s = append(*s, t)
		return
	}

	if root.Left != nil {
		btpHelper(root.Left, s, t + "->")
	}
	if root.Right != nil {
		btpHelper(root.Right, s, t + "->")
	}
}