package main

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	var root TreeNode
	m := length/2
	root.Val = nums[m]
	root.Left = sortedArrayToBST(nums[:m])
	root.Right = sortedArrayToBST(nums[m+1:])
	return &root
}