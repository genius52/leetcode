package number

func dp_peopleAwareOfSecret(cur int,n int,delay int, forget int,memo []int)int{
	if cur + delay > n{
		return 1
	}
	if memo[cur] != 0{
		return memo[cur]
	}
	var MOD int = 1e9 + 7
	var res int = 0
	if cur + forget > n{
		res = 1
	}
	for i := cur + delay;i <= n && i < cur + forget;i++{
		res += dp_peopleAwareOfSecret(i,n,delay,forget,memo)
		res %= MOD
	}
	memo[cur] = res
	return res
}

func PeopleAwareOfSecret(n int, delay int, forget int) int {
	var memo []int = make([]int,n + 1)
	return dp_peopleAwareOfSecret(1,n,delay,forget,memo)
}

func PeopleAwareOfSecret2(n int, delay int, forget int) int {
	var dp []int = make([]int,n + 1)
	dp[1] = 1
	var MOD int = 1e9 + 7
	var res int = 0
	for i := 1;i <= n;i++{
		if i + forget > n{
			res += dp[i]
		}
		for j := i + delay;j < i + forget && j <= n;j++{
			dp[j] += dp[i]
			dp[j] %= MOD
		}
	}
	return res
}