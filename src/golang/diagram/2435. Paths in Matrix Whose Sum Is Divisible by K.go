package diagram

//1,5,3,7,3,2,3,5
func NumberOfPaths(grid [][]int, k int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var dp [][][]int = make([][][]int,rows)
	var MOD int = 1e9 + 7
	for i := 0;i < rows;i++{
		dp[i] = make([][]int,columns)
		for j := 0;j < columns;j++{
			dp[i][j] = make([]int,k)
		}
	}
	var sum int = grid[0][0]
	dp[0][0][sum % k] = 1
	for i := 1;i < rows;i++{
		sum += grid[i][0]
		dp[i][0][sum % k] = 1
	}
	sum = grid[0][0]
	for j := 1;j < columns;j++{
		sum += grid[0][j]
		dp[0][j][sum % k] = 1
	}
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			for m := 0;m < k;m++{
				if dp[i - 1][j][m] != 0{
					dp[i][j][(m + grid[i][j]) % k] += dp[i - 1][j][m]
				}
				if dp[i][j - 1][m] != 0{
					dp[i][j][(m + grid[i][j]) % k] += dp[i][j - 1][m]
				}
				dp[i][j][(m + grid[i][j]) % k] %= MOD
			}
		}
	}
	return dp[rows - 1][columns - 1][0]
}