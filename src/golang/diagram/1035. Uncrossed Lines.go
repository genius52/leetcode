package diagram

import "math"

//1 2 3
//2 3 1
func dfs_maxUncrossedLines(A []int, B []int,len_a int,len_b int,pos_a int,pos_b int,memo [][]int)int{
	if pos_a >= len_a || pos_b >= len_b{
		return 0
	}
	if memo[pos_a][pos_b] != -1{
		return memo[pos_a][pos_b]
	}
	var res int = 0
	if A[pos_a] == B[pos_b]{
		res = 1 + dfs_maxUncrossedLines(A,B,len_a,len_b,pos_a + 1,pos_b + 1,memo)
	}
	skip_a := dfs_maxUncrossedLines(A,B,len_a,len_b,pos_a + 1,pos_b,memo)
	skip_b := dfs_maxUncrossedLines(A,B,len_a,len_b,pos_a,pos_b + 1,memo)
	if skip_a > res{
		res = skip_a
	}
	if skip_b > res{
		res = skip_b
	}
	memo[pos_a][pos_b] = res
	return res
}

func MaxUncrossedLines2(A []int, B []int) int {
	var len_a int = len(A)
	var len_b int = len(B)
	var memo [][]int = make([][]int,len_a)
	for i := 0;i < len_a;i++{
		memo[i] = make([]int,len_b)
		for j := 0;j < len_b;j++{
			memo[i][j] = -1
		}
	}
	return dfs_maxUncrossedLines(A,B,len_a,len_b,0,0,memo)
}

func MaxUncrossedLines(A []int, B []int) int {
	var len_a int = len(A)
	var len_b int = len(B)
	var dp [][]int = make([][]int,len_a)
	for i := 0;i < len_a;i++{
		dp[i] = make([]int,len_b)
	}
	for i := 0;i < len_a;i++{
		for j := 0;j < len_b;j++{
			if A[i] == B[j]{
				if i == 0 || j == 0{
					dp[i][j] = 1
				}else {
					if dp[i][j] < (1 + dp[i - 1][j - 1]){
						dp[i][j] = 1 + dp[i - 1][j - 1]
					}
				}
			}else{
				if i == 0 && j == 0{
					dp[i][j] = 0
				}else if i == 0{
					dp[i][j] = dp[i][j - 1]
				}else if j == 0{
					dp[i][j] = dp[i - 1][j]
				}else{
					dp[i][j] = int(math.Max(float64(dp[i - 1][j]),float64(dp[i][j - 1])))
					if dp[i - 1][j - 1] > dp[i][j]{
						dp[i][j] = dp[i - 1][j - 1]
					}
				}
			}
		}
	}
	return dp[len_a - 1][len_b - 1]
}