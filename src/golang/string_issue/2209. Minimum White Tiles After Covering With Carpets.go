package string_issue

func dp_minimumWhiteTiles(floor string,l int,idx int,numCarpets int,carpetLen int,memo [][]int)int{
	if idx >= l{
		return 0
	}
	if memo[idx][numCarpets] != 2147483647{
		return memo[idx][numCarpets]
	}
	var res int = 2147483647
	if floor[idx] == '0'{
		res = dp_minimumWhiteTiles(floor,l,idx + 1,numCarpets,carpetLen,memo)
	}else{
		res = 1 + dp_minimumWhiteTiles(floor,l,idx + 1,numCarpets,carpetLen,memo)
		if numCarpets > 0{
			cover := dp_minimumWhiteTiles(floor,l,idx + carpetLen,numCarpets - 1,carpetLen,memo)
			if cover < res{
				res = cover
			}
		}
	}
	memo[idx][numCarpets] = res
	return res
}

func MinimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	var l int = len(floor)
	var memo [][]int = make([][]int,l + 1)//memo[i][j]: 在i结尾使用了j条地毯时，最少的白色瓷砖
	for i := 0;i <= l;i++{
		memo[i] = make([]int,numCarpets + 1)
		for j := 0;j <= numCarpets;j++{
			memo[i][j] = 2147483647
		}
	}
	return dp_minimumWhiteTiles(floor,l,0,numCarpets,carpetLen,memo)
}