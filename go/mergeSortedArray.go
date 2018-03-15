package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	cur := m + n - 1
	p1 := m - 1
	p2 := n - 1
	for {
		if p1 < 0 || p2 < 0 {
			break
		}

		if nums1[p1] >= nums2[p2] {
			nums1[cur] = nums1[p1]
			p1--
		} else {
			nums1[cur] = nums2[p2]
			p2--
		}

		cur--
	}

	for {
		if p2 < 0 {
			break
		}
		nums1[cur] = nums2[p2]
		p2--
		cur--
	}
}
