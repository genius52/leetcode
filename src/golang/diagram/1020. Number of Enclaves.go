package diagram

func dfs_numEnclaves(A [][]int,r int,c int,rows int,columns int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if A[r][c] == 0{
		return
	}
	A[r][c] = 0
	dfs_numEnclaves(A,r - 1,c,rows,columns)
	dfs_numEnclaves(A,r + 1,c,rows,columns)
	dfs_numEnclaves(A,r,c - 1,rows,columns)
	dfs_numEnclaves(A,r,c + 1,rows,columns)
}

func numEnclaves(A [][]int) int {
	var rows int = len(A)
	var columns int = len(A[0])
	for i := 0;i < rows;i++{
		if A[i][0] != 0 {
			dfs_numEnclaves(A,i,0,rows,columns)
		}
		if A[i][columns - 1] != 0{
			dfs_numEnclaves(A,i,columns - 1,rows,columns)
		}
	}
	for j := 0;j < columns;j++{
		if A[0][j] != 0 {
			dfs_numEnclaves(A,0,j,rows,columns)
		}
		if A[rows - 1][j] != 0{
			dfs_numEnclaves(A,rows - 1,j,rows,columns)
		}
	}
	var res int = 0
	for i := 1;i < rows - 1;i++{
		for j := 1;j < columns;j++{
			if A[i][j] == 1{
				res++
			}
		}
	}
	return res
}