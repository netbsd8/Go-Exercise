package main

func convert(s string, numRows int) string {
	length := len(s)
	ret := ""

	if length == 0 {
		return ret
	}

	if numRows == 1 {
		return s
	}

	row := make([]string, numRows)
	curRowNum := 0
	direction := 1
	for i := 0; i < length; i++ {
		row[curRowNum] = row[curRowNum] + string(s[i])

		curRowNum = curRowNum + direction

		if curRowNum == numRows {
			curRowNum = numRows - 2
			direction = -1
		}

		if curRowNum < 0 {
			curRowNum = 1
			direction = 1
		}
	}

	for i := 0; i < numRows; i++ {
		ret = ret + row[i]
	}
	return ret
}
