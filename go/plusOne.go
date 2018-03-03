package main

func plusOne(digits []int) []int {
	ret := make([]int, 0)
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		value := digits[i] + carry
		cur := value % 10
		carry = value / 10

		ret = insert(ret, 0, cur)
	}
	if carry > 0 {
		ret = insert(ret, 0, 1)
	}

	return ret
}

func insert(slice []int, index, value int) []int {
	slc := make([]int, len(slice)+1)
	copy(slc[index+1:], slice[index:])
	slc[index] = value
	return slc
}
