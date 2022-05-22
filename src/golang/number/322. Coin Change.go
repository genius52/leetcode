package number

import (
	"math"
	"sort"
)

func CoinChange(coins []int, amount int) int {
	var dp []int = make([]int,amount + 1)//dp[i] i元所需的最少硬币数
	var l int = len(coins)
	sort.Ints(coins)
	for i := 0;i < l;i++{
		if coins[i] <= amount{
			dp[coins[i]] = 1
		}
	}
	for i := 1;i <= amount;i++{
		if dp[i] == 1{
			continue
		}
		dp[i] = 2147483647
		for j := 0;j < l;j++{
			if (i - coins[j]) >= 0 && dp[i - coins[j]] != 0{
				dp[i] = min_int(dp[i],1 + dp[i - coins[j]])
			}
		}
	}
	if dp[amount] == 2147483647{
		return -1
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	var dp []int = make([]int,amount + 1)
	for i := 0;i < len(dp);i++{
		dp[i] = -1
	}
	dp[0] = 0
	var min_coins int = math.MaxInt32
	for i := 1;i <= amount;i++{
		min_coins = math.MaxInt32
		for _,c := range coins{
			if c > i{
				continue
			}
			if dp[i - c] != - 1{
				min_coins = min_int(1 + dp[i-c],min_coins)
			}
		}
		if min_coins != math.MaxInt32{
			dp[i] = min_coins
		}
	}
	return dp[amount]
}

func dfs_coinChange(coins []int, amount int,record map[int]int)int{
	if amount == 0{
		return 0
	}
	if _,ok := record[amount];ok{
		return record[amount]
	}
	steps := math.MaxInt32
	for _,c := range coins{
		if c > amount{
			continue
		}
		res := dfs_coinChange(coins,amount - c,record)
		if res != -1{
			steps = min_int(res + 1,steps)
		}
	}
	if steps == math.MaxInt32{
		return -1
	}else{
		record[amount] = steps
		return steps
	}
}

func coinChange2(coins []int, amount int)int{
	if amount == 0{
		return 0
	}

	var coin_record map[int]int = make(map[int]int)
	return dfs_coinChange(coins,amount,coin_record)
}