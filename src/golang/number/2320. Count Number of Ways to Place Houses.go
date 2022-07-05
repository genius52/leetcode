package number

func dp_countHousePlacements(n int,idx int,up bool,down bool,memo [][2][2]int)int{
	if idx == n{
		return 1
	}
	var idx1 int = 0
	var idx2 int = 0
	if up{
		idx1 = 1
	}
	if down{
		idx2 = 1
	}
	if memo[idx][idx1][idx2] != 0{
		return memo[idx][idx1][idx2]
	}
	var res int = 0
	if !up{
		res += dp_countHousePlacements(n,idx + 1,true,false,memo)
	}
	if !down{
		res += dp_countHousePlacements(n,idx + 1,false,true,memo)
	}
	if !up && !down{
		res += dp_countHousePlacements(n,idx + 1,true,true,memo)
	}
	res += dp_countHousePlacements(n,idx + 1,false,false,memo)
	res %= 1e9 + 7
	memo[idx][idx1][idx2] = res
	return res
}

func countHousePlacements(n int) int {
	//var dp [][2]int = make([][2]int,n + 1) //dp[i][0]:  在位置i,
	var memo [][2][2]int = make([][2][2]int,n + 1)
	return dp_countHousePlacements(n,0,false,false,memo)
}