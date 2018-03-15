package main

func generatePascalI(numRows int) [][]int {
	ret := [][]int{}

	if numRows == 0 {
		return ret
	}

	prevRow := make([]int, 0)
	for i:=0; i<numRows; i++ {
		curRow := make([]int, i+1)
		for j:=0; j<=i; j++ {
			if i >= 2 && (j>0 && j < i) {
				curRow[j] = prevRow[j-1] + prevRow[j] 
			}
		}
		curRow[0] = 1
		curRow[i] = 1
		prevRow = curRow
		ret = append(ret, curRow)
	}

	return ret
}

/*
func main() {
	ans := generatePascalI(5)
	fmt.Println(ans)
}
*/