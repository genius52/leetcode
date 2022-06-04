package array

func dfs_hasValidPath(grid [][]byte,rows int,columns int,r int,c int,left int,right int,memo map[int]map[int]bool)bool{
	if r >= rows || c >= columns{
		return false
	}
	if grid[r][c] == '('{
		left++
	}else{
		right++
	}
	if right > left{
		return false
	}
	if r == rows - 1 && c == columns - 1{
		return left == right
	}
	k1 := r * 100 + c
	k2 := left * 100 + right
	if _,ok1 := memo[k1];ok1{
		if _,ok2 := memo[k1][k2];ok2{
			return false
		}
	}else{
		memo[k1] = make(map[int]bool)
	}
	memo[k1][k2] = true
	return dfs_hasValidPath(grid,rows,columns,r + 1,c,left,right,memo) || dfs_hasValidPath(grid,rows,columns,r,c + 1,left,right,memo)
}

func hasValidPath(grid [][]byte) bool {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var memo map[int]map[int]bool = make(map[int]map[int]bool)
	return dfs_hasValidPath(grid,rows,columns,0,0,0,0,memo)
}