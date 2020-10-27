package number

func ChampagneTower(poured int, query_row int, query_glass int) float64{
	if poured <= 0 {
		return 0.0
	}
	var dp [][]float64 = make([][]float64,101)
	for i := 0;i < 101;i++{
		dp[i] = make([]float64,101)
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

//func champagneTower(poured int, query_row int, query_glass int) float64 {
	//if poured <= 0 || query_row > poured / 2{
	//	return 0.0
	//}
	//var total float64 = float64((2 + query_row)* (query_row + 1)/ 2)
	//if int(total) <= poured{
	//	return 1.0
	//}
	//var total2 float64 = float64((1 + query_row)* (query_row)/ 2)
	//var rest float64 = float64(poured) - total2
	//if query_glass == 0 || query_glass == query_row{
	//	return 0.0
	//}
	//return rest/float64(query_row - 1)
//}