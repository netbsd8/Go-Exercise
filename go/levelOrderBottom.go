package main

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ret := make([][]int, 0)

	q := make(chan *TreeNode)
	q <- root
	
	for {
		if len(q) == 0 {
			break
		}

		curSlice := make([]int, 0)
		length := len(q)
		for i:=0; i<length; i++ {
			cur :=<- q
			if cur.Left != nil {
				q <- cur.Left
			}
			if cur.Right != nil {
				q <- cur.Right
			}
			curSlice = append(curSlice, cur.Val)
		}
		ret = append(ret, curSlice)
	}

	return reverseResult(ret)
}

func reverseResult(data [][]int) [][]int {
	r := len(data) - 1
	l := 0

	for {
		if l >= r {
			break
		}
		data[l], data[r] = data[r], data[l]
		l++
		r--
	}

	return data
}

func levelOrderBottom3(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	levelOrderHelper(&ret, 0, root)

	return reverseResult(ret)
}

func levelOrderHelper(ret *[][]int, level int, root *TreeNode) {
	if ret == nil {
		return
	}

	if len(*ret) == level {
		*ret = append(*ret, []int{})
	} 
	(*ret)[level] = append((*ret)[level], root.Val)
	if root.Left != nil {
		levelOrderHelper(ret, level+1, root.Left)
	}	
	if root.Right != nil {
		levelOrderHelper(ret, level+1, root.Right)
	}
}