package zizagconversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]rune, numRows)
	curRow := 0
	goDown := false
	for _, c := range s {
		rows[curRow] = append(rows[curRow], c)
		if curRow == numRows-1 || curRow == 0 {
			goDown = !goDown
		}
		if goDown {
			curRow++
		} else {
			curRow--
		}
	}
	res := make([]rune, 0, len(s))
	for _, row := range rows {
		res = append(res, row...)
	}
	return string(res)
}
