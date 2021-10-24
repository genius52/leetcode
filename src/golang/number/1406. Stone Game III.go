package number

//输入：values = [1,2,3,7]
//输出："Bob"
//解释：Alice 总是会输，她的最佳选择是拿走前三堆，得分变成 6 。但是 Bob 的得分为 7，Bob 获胜。

func dp_stoneGameIII(stoneValue []int,l int,pos int,Alice bool,memo *[][2]int)int{
	if pos >= l{
		return 0
	}
	if Alice{
		if (*memo)[pos][1] != -1{
			return (*memo)[pos][1]
		}
	}else{
		if (*memo)[pos][0] != -1{
			return (*memo)[pos][0]
		}
	}
	if Alice{
		res := -2147483648
		var sum int = 0
		for i := pos;i < l && i < pos + 3;i++{
			sum += stoneValue[i]
			res = max_int(res,sum + dp_stoneGameIII(stoneValue,l,i + 1,!Alice,memo))
		}
		(*memo)[pos][1] = res
		return res
	}else{
		res := 2147483647
		for i := pos;i < l && i < pos + 3;i++{
			res = min_int(res,dp_stoneGameIII(stoneValue,l,i + 1,!Alice,memo))
		}
		(*memo)[pos][0] = res
		return res
	}
}

func StoneGameIII(stoneValue []int) string {
	var sum int = 0
	var l int = len(stoneValue)
	for i := 0;i < l;i++{
		sum += stoneValue[i]
	}
	var memo [][2]int = make([][2]int,l)
	for i := 0;i < l;i++{
		memo[i][0] = -1
		memo[i][1] = -1
	}
	res := dp_stoneGameIII(stoneValue,l,0,true,&memo)
	rest := sum - res
	if res > rest{
		return "Alice"
	}else if res < rest{
		return "Bob"
	}else{
		return "Tie"
	}
}