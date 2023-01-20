package array

func RangeAddQueries(n int, queries [][]int) [][]int {
	var res [][]int = make([][]int,n)
	var prefix [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		res[i] = make([]int,n)
		prefix[i] = make([]int,n)
	}
	for _,query := range queries{
		row1 := query[0]
		col1 := query[1]
		row2 := query[2]
		col2 := query[3]
		if row1 > 0{
			prefix[row1 - 1][col2]--
		}
		if col1 > 0{
			prefix[row2][col1 - 1]--
		}
		if row1 > 0 && col1 > 0{
			prefix[row1 - 1][col1 - 1]++
		}
		prefix[row2][col2]++
	}
	for i := n - 1;i >= 0;i--{
		for j := n - 1;j >= 0;j--{
			if j + 1 < n{
				prefix[i][j] += prefix[i][j + 1]
			}
		}
	}
	for j := n - 1;j >= 0;j--{
		for i := n - 1;i >= 0;i--{
			if i + 1 < n{
				prefix[i][j] += prefix[i + 1][j]
			}
		}
	}
	for i := 0;i < n;i++{
		for j := 0;j < n;j++{
			res[i][j] = prefix[i][j]
		}
	}
	return res
}