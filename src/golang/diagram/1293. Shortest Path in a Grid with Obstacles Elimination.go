package diagram

func dfs_shortestPath(grid *[][]int,rows int,columns int,r int,c int,k int,memo *[][][]int)(bool,int){
	if r < 0 || r >= rows || c < 0 || c >= columns ||  (*grid)[r][c] == 2{
		return false,-1
	}
	if (*memo)[k][r][c] != -1{
		if (*memo)[k][r][c] == -2{
			return false,-1
		}else{
			return true,(*memo)[k][r][c]
		}
	}
	if r == (rows - 1) && c == (columns - 1){
		return true,0
	}
	if (*grid)[r][c] == 1{
		if k == 0{
			return false,-1
		}
		k--
	}
	old := (*grid)[r][c]
	(*grid)[r][c] = 2
	var res bool = false
	var steps int = 2147483647
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,dir := range dirs{
		next_r := r + dir[0]
		next_c := c + dir[1]
		cur_res,cur_steps := dfs_shortestPath(grid,rows,columns,next_r,next_c,k,memo)
		if cur_res{
			res = true
			steps = min_int(steps,1 + cur_steps)
		}
	}
	(*grid)[r][c] = old
	if res{
		(*memo)[k][r][c] = steps
		return true,steps
	}else{
		(*memo)[k][r][c] = -2
		return false,-2
	}
}

func ShortestPath(grid [][]int, k int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var memo [][][]int = make([][][]int,k + 1)
	for m := 0;m <= k;m++{
		memo[m] = make([][]int,rows)
		for i := 0;i < rows;i++{
			memo[m][i] = make([]int,columns)
			for j := 0;j < columns;j++{
				memo[m][i][j] = -1
			}
		}
	}
	ret,steps := dfs_shortestPath(&grid,rows,columns,0,0,k,&memo)
	if ret{
		return steps
	}else{
		return -1
	}
}