package array

import "math"

//64
//Input:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//Output: 7
//Explanation: Because the path 1→3→1→1→1 minimizes the sum.
func dp_minpathsum(grid [][]int,row int,col int,target_row int,target_col int,visited [][]int)(int,bool){
	if row < 0|| row > target_row || col < 0 || col > target_col {
		return 0,false
	}
	if visited[row][col] != 0{
		return visited[row][col],true
	}
	if row == target_row && col == target_col{
		visited[row][col] = grid[row][col]
		return grid[row][col],true
	}
	var sum int = math.MaxInt32
	var res bool = false
	if val,ok := dp_minpathsum(grid,row + 1,col,target_row,target_col,visited);ok{
		sum = min_int_number(grid[row][col] + val,sum)
		res = ok
	}
	if val,ok := dp_minpathsum(grid,row,col + 1,target_row,target_col,visited);ok{
		sum = min_int_number(grid[row][col] + val,sum)
		res = ok
	}
	if res{
		visited[row][col] = sum
	}
	return sum,res
}

func minPathSum(grid [][]int) int {
	var visited [][]int = make([][]int,len(grid))
	for i := 0;i < len(grid);i++{
		visited[i] = make([]int,len(grid[0]))
	}
	target_row := len(grid) - 1
	target_col := len(grid[0]) - 1
	val,_ :=  dp_minpathsum(grid,0,0,target_row,target_col,visited)
	return val
}

func MinPathSum2(grid [][]int) int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	var dp [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
	}
	dp[0][0] = grid[0][0]
	for i := 1;i < rows;i++{
		dp[i][0] = grid[i][0] + dp[i - 1][0]
	}
	for j := 1;j < columns;j++{
		dp[0][j] = grid[0][j] + dp[0][j - 1]
	}
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			if dp[i - 1][j] < dp[i][j - 1]{
				dp[i][j] = grid[i][j] + dp[i - 1][j]
			}else{
				dp[i][j] = grid[i][j] + dp[i][j - 1]
			}
		}
	}
	return dp[rows - 1][columns - 1]
}