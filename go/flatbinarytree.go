package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	helper(root)
}

func helper(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	leftLastNode := helper(node.Left)
	rightLastNode := helper(node.Right)

	if leftLastNode != nil && rightLastNode != nil {
		leftLastNode.Right = node.Right
		node.Right = node.Left
		node.Left = nil
		return rightLastNode
	}

	if leftLastNode != nil {
		return leftLastNode
	}

	if rightLastNode != nil {
		return rightLastNode
	}

	return node
}

//iterative
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	cur := root
	for {
		if cur == nil {
			return
		}

		if cur.Left == nil {
			cur = cur.Right
			continue
		}

		prev := cur
		cur = cur.Left
		for {
			if cur.Right == nil {
				break
			}
			cur = cur.Right
		}
		cur.Right = prev.Right
		prev.Right = prev.Left
		prev.Left = nil
		cur = cur.Right
	}
}
