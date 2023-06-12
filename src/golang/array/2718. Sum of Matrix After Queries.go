package array

func matrixSumQueries(n int, queries [][]int) int64 {
	var l int = len(queries)
	var res int64 = 0
	var visited_row []bool = make([]bool, n)
	var visited_col []bool = make([]bool, n)
	for i := l - 1; i >= 0; i-- {
		t := queries[i][0]
		idx := queries[i][1]
		val := queries[i][2]
		if t == 0 {
			if visited_row[idx] {
				continue
			}
			visited_row[idx] = true
			for i := 0; i < n; i++ {
				if !visited_col[i] {
					res += int64(val)
				}
			}
		} else if t == 1 {
			if visited_col[idx] {
				continue
			}
			visited_col[idx] = true
			for i := 0; i < n; i++ {
				if !visited_row[i] {
					res += int64(val)
				}
			}
		}
	}
	return res
}
