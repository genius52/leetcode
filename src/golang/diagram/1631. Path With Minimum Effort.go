package diagram

import "math"

func dfs_minimumEffortPath(heights [][]int,row int,col int,last_height int,target_row int,target_col int,memo [][]int,visited [][]bool)int{
	if row < 0 || row > target_row || col < 0 || col > target_col{
		return math.MaxInt32
	}
	if visited[row][col]{
		return math.MaxInt32
	}
	cur_diff := int(math.Abs(float64(heights[row][col]) - float64(last_height)))
	if row == target_row && col == target_col{
		return cur_diff
	}
	if memo[row][col] != -1{
		return int(math.Max(float64(cur_diff),float64(memo[row][col])))
	}
	visited[row][col] = true
	var sub_diff int = math.MaxInt32
	if row > 0{
		r1 := dfs_minimumEffortPath(heights,row - 1,col,heights[row][col],target_row,target_col,memo,visited)
		if r1 < sub_diff{
			sub_diff = r1
		}
	}
	if row < target_row{
		r2 := dfs_minimumEffortPath(heights,row + 1,col,heights[row][col],target_row,target_col,memo,visited)
		if r2 < sub_diff{
			sub_diff = r2
		}
	}
	if col > 0{
		c1 := dfs_minimumEffortPath(heights,row,col - 1,heights[row][col],target_row,target_col,memo,visited)
		if c1 < sub_diff{
			sub_diff = c1
		}
	}
	if col < target_col{
		c2 := dfs_minimumEffortPath(heights,row,col + 1,heights[row][col],target_row,target_col,memo,visited)
		if c2 < sub_diff{
			sub_diff = c2
		}
	}
	cur_diff = int(math.Max(float64(cur_diff),float64(sub_diff)))
	memo[row][col] = sub_diff
	visited[row][col] = false
	return cur_diff
}

func MinimumEffortPath(heights [][]int) int {
	var rows int = len(heights)
	var columns int = len(heights[0])
	if rows == 1 && columns == 1{
		return 0
	}
	var memo [][]int = make([][]int,rows)
	var visited [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		memo[i] = make([]int,columns)
		visited[i] = make([]bool,columns)
		for j := 0;j < columns;j++{
			memo[i][j] = -1
		}
	}
	visited[0][0] = true
	return int(math.Min(float64(dfs_minimumEffortPath(heights,1,0,heights[0][0],rows - 1,columns - 1,memo,visited)),
	float64(dfs_minimumEffortPath(heights,0,1,heights[0][0],rows - 1,columns - 1,memo,visited))))
}