package array

//Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//Output: 2
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	var rows int = len(obstacleGrid)
	var columns int = len(obstacleGrid[0])
	var dp [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		if obstacleGrid[i][0] == 1{
			break
		}
		dp[i][0] = 1
	}
	for i := 0;i < columns;i++{
		if obstacleGrid[0][i] == 1{
			break
		}
		dp[0][i] = 1
	}
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			if obstacleGrid[i][j] == 1{
				dp[i][j] = 0
			}else{
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}
	return dp[rows - 1][columns - 1]
}