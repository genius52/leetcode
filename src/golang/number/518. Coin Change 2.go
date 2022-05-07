package number

import (
	"sort"
)

//Input: amount = 5, coins = [1, 2, 5]
//Output: 4
//Explanation: there are four ways to make up the amount:
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
func dfs_change(amount int,coins []int,l int,pos int,memo [][]int)int{
	if amount < 0{
		return 0
	}
	if amount == 0{
		return 1
	}
	if pos >= l{
		return 0
	}
	if memo[pos][amount] != -1{
		return memo[pos][amount]
	}
	var total int = 0
	for j := 0;coins[pos] * j <= amount;j++{
		total += dfs_change(amount - coins[pos] * j,coins,l,pos + 1,memo)
	}
	memo[pos][amount] = total
	return total
}

func Change(amount int, coins []int) int {
	sort.Ints(coins)
	var l int = len(coins)
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,amount + 1)
		for j := 0;j <= amount;j++{
			memo[i][j] = -1
		}
	}
	return dfs_change(amount,coins,l,0,memo)
}

func Change2(amount int, coins []int) int {
	sort.Ints(coins)
	var l int = len(coins)
	var dp [][]int = make([][]int,l + 1)//dp[i][j] = the kinds form to j by using (i - 1)th element
	for i := 0;i <= l;i++{
		dp[i] = make([]int,amount + 1)
		dp[i][0] = 1
	}
	dp[0][0] = 1
	for i := 1;i <= l;i++{
		for j := 1;j <= amount;j++{
			if j >= coins[i - 1]{
				dp[i][j] += dp[i][j - coins[i - 1]] //
			}
			dp[i][j] += dp[i-1][j]//do not choose (i)th element
		}
	}
	return dp[l][amount]
}