package number

//Serve 100 ml of soup A and 0 ml of soup B
//Serve 75 ml of soup A and 25 ml of soup B
//Serve 50 ml of soup A and 50 ml of soup B
//Serve 25 ml of soup A and 75 ml of soup B

//Input: n = 50
//Output: 0.625
//0.25 * (1 + 1 + 0.5 + 0) = 0.625.
func dfs_soupServings(a int,b int,memo map[int]map[int]float64)float64{
	if a <= 0 && b <= 0{
		return 0.5
	}else if a <= 0{
		return 1.0
	}else if b <= 0{
		return 0.0
	}else{
		if _,ok1 := memo[a];ok1{
			if _,ok2 := memo[a][b];ok2{
				return memo[a][b]
			}
		}else{
			memo[a] = make(map[int]float64)
		}
		res := 0.25 * dfs_soupServings(a - 100,b,memo) + 0.25 * dfs_soupServings(a - 75,b - 25,memo) +
			0.25 * dfs_soupServings(a - 50,b - 50,memo) + 0.25 * dfs_soupServings(a - 25,b - 75,memo)
		memo[a][b] = res
		return res
	}
}

func SoupServings(N int) float64 {
	if N > 4800{
		return 1
	}
	var memo map[int]map[int]float64 = make(map[int]map[int]float64)
	return dfs_soupServings(N,N,memo)
}