package number

func ChampagneTower(poured int, query_row int, query_glass int) float64{
	if poured <= 0 {
		return 0.0
	}
	var dp [][]float64 = make([][]float64,query_row + 2)
	for i := 0;i <= query_row + 1;i++{
		dp[i] = make([]float64,i + 1)
	}
	dp[0][0] = float64(poured)
	for r := 0;r <= query_row;r++{
		for c := 0;c <= r;c++{
			if dp[r][c] >= 1.0{
				dp[r + 1][c] += (dp[r][c] - 1.0)/2
				dp[r + 1][c + 1] += (dp[r][c] - 1.0)/2
				dp[r][c] = 1.0
			}
		}
	}
	return dp[query_row][query_glass]
}