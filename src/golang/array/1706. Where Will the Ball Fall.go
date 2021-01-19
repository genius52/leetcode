package array

//Input: grid = [[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]
//Output: [1,-1,-1,-1,-1]
const (
	left int = 1
	right int = 2
	down int = 3
)

func dfs_findBall(grid [][]int,rows int,columns int,r int,c int,prev_dir int)int{
	if r >= rows{
		return c
	}
	if c < 0 || c >= columns{
		return -1
	}
	if prev_dir == down{
		if grid[r][c] == 1{
			return dfs_findBall(grid,rows,columns,r,c + 1,right)
		}else{
			return dfs_findBall(grid,rows,columns,r,c - 1,left)
		}
	}else if prev_dir == left{
		if grid[r][c] == 1{
			return -1
		}else{
			return dfs_findBall(grid,rows,columns,r + 1,c,down)
		}
	}else if prev_dir == right{
		if grid[r][c] == -1{
			return -1
		}else{
			return dfs_findBall(grid,rows,columns,r + 1,c,down)
		}
	}
	return -1
}

func FindBall(grid [][]int) []int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res []int = make([]int,columns)
	for i := 0;i < columns;i++{
		ret := dfs_findBall(grid,rows,columns,0,i,down)
		res[i] = ret
	}
	return res
}