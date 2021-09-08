package array

//n = 2, m = 3, indices = [[0,1],[1,1]]
func OddCells(n int, m int, indices [][]int) int{
	var res int = 0
	var rows_odd []bool = make([]bool,n)
	var columns_odd []bool = make([]bool,m)
	for _,index := range indices{
		rows_odd[index[0]] = !rows_odd[index[0]]
		columns_odd[index[1]] = !columns_odd[index[1]]
	}
	for i := 0;i < n;i++{
		for j := 0;j < m;j++{
			if (rows_odd[i] && !columns_odd[j]) || (!rows_odd[i] && columns_odd[j]){
				res++
			}
		}
	}
	return res
}

func oddCells(n int, m int, indices [][]int) int {
	var res int = 0
	var matrix [][]int = make([][]int,n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, m)
	}
	for i:= 0; i < len(indices);i++{
		row := indices[i][0]
		for c:= 0; c < m;c++{
			matrix[row][c]++
		}
		column := indices[i][1]
		for r := 0;r < n;r++{
			matrix[r][column]++
		}
	}
	for i := 0; i < n;i++{
		for j := 0;j < m;j++{
			if matrix[i][j]%2 == 1{
				res++
			}
		}
	}
	return res
}