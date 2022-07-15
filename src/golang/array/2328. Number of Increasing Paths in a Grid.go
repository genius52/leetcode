package array

func dp_countPaths(grid [][]int,rows int,columns int,r int,c int,pre int,memo [][]int)int{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return 0
	}
	if grid[r][c] <= pre{
		return 0
	}
	if memo[r][c] != 0{
		return memo[r][c]
	}
	var res int = 1
	var MOD int = 1e9 + 7
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,dir := range dirs{
		next_r := r + dir[0]
		next_c := c + dir[1]
		res += dp_countPaths(grid,rows,columns,next_r,next_c,grid[r][c],memo)
		res %= MOD
	}
	memo[r][c] = res
	return res
}

func CountPaths(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var memo [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		memo[i] = make([]int,columns)
	}
	var res int = 0
	var MOD int = 1e9 + 7
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			res += dp_countPaths(grid,rows,columns,i,j,-1,memo)
			res %= MOD
		}
	}
	return res
}