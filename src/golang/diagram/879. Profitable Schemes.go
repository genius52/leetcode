package diagram

func dp_profitableSchemes(group []int, profit []int,l int,index int,available int,require_profit int,memo *[101][101][101]int)int{
	if index >= l{
		return 0
	}
	var res int = 0
	if require_profit < 0{
		require_profit = 0
	}
	if memo[index][available][require_profit] != 0{
		return memo[index][available][require_profit]
	}
	if group[index] <= available{
		if (require_profit - profit[index]) <= 0{
			res++
		}
		//choose current
		res += dp_profitableSchemes(group,profit,l,index + 1,available - group[index],require_profit - profit[index],memo)
	}
	//skip current
	res += dp_profitableSchemes(group,profit,l,index + 1,available,require_profit,memo);
	memo[index][available][require_profit] = res
 	return res % 1000000007
}

func ProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	var l int = len(group)
	var memo [101][101][101]int
	//for i := 0;i < l;i++{
	//	memo[i] = make([][]int,n + 1)
	//	for j := 0;j <= n;j++{
	//		memo[i][j] = make([]int,minProfit + 1)
	//	}
	//}
	res := dp_profitableSchemes(group,profit,l,0,n,minProfit,&memo)
	if minProfit == 0{
		res++
	}
	return res
}