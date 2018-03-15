package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmHelper(root.Left, root.Right)
}

func symmHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if (left == nil && right != nil) || (left != nil && right == nil) || (left.Val != right.Val) {
		return false
	}

	return symmHelper(left.Left, right.Right) && symmHelper(left.Right, right.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q1 := make(chan *TreeNode)
	q2 := make(chan *TreeNode)

	q1 <- root.Left
	q2 <- root.Right
	var c1 *TreeNode
	var c2 *TreeNode

	for {
		if len(q1) ==0 || len(q2)==0 {
			break
		}
		c1 =<- q1
		c2 =<- q2

		if (c1 == nil && c2 != nil) || (c1 != nil && c2 == nil) {
			return false
		}
		if c1.Val != c2.Val {
			return false
		}
		if (c1 == nil && c2 == nil) {
			continue
		}

		q1 <- c1.Left
		q1 <- c1.Right
		q2 <- c2.Right
		q2 <- c2.Left
	}

	return len(q1) == 0 && len(q2) == 0
}