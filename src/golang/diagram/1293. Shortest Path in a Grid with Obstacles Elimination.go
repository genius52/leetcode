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
	res1,steps1 := dfs_shortestPath(grid,rows,columns,r - 1,c,k,memo)
	res2,steps2 := dfs_shortestPath(grid,rows,columns,r + 1,c,k,memo)
	res3,steps3 := dfs_shortestPath(grid,rows,columns,r,c - 1,k,memo)
	res4,steps4 := dfs_shortestPath(grid,rows,columns,r,c + 1,k,memo)
	(*grid)[r][c] = old
	res = res1 || res2 || res3 || res4
	if res{
		if steps1 >= 0{
			steps = min_int(steps,1 + steps1)
		}
		if steps2 >= 0{
			steps = min_int(steps,1 + steps2)
		}
		if steps3 >= 0{
			steps = min_int(steps,1 + steps3)
		}
		if steps4 >= 0{
			steps = min_int(steps,1 + steps4)
		}
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