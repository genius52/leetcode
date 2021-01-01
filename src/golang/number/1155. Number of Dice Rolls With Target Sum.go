package number

//Input: d = 2, f = 6, target = 7
//Output: 6
//Explanation:
//You throw two dice, each with 6 faces.  There are 6 ways to get a sum of 7:
//1+6, 2+5, 3+4, 4+3, +52, 6+1.
func dp_numRollsToTarget(d int,f int,target int,memo map[string]int)int{
	if d == 0 {
		if target == 0{
			return 1
		}
		return 0
	}
	if d * f < target || d > target || target < 0{
		return 0
	}
	if d == 1{
		return 1
	}
	key := string(d) + "," + string(target)
	if n,ok := memo[key];ok{
		return n
	}
	var total int = 0
	for i := 1;i <= f;i++{
		if i > target{
			break
		}
		total += dp_numRollsToTarget(d - 1,f,target - i,memo)
		total %= 1000000007
	}
	memo[key] = total
	return total
}

func NumRollsToTarget(d int, f int, target int) int {
	var memo map[string]int = make(map[string]int)
	return dp_numRollsToTarget(d,f,target,memo)
}

func NumRollsToTarget2(d int, f int, target int)int{
	var dp [][]int = make([][]int,d + 1)//dp[]
	var m int = target
	if f > target{
		m = f
	}
	for i := 0;i <= d;i++{
		dp[i] = make([]int,m + 1)
	}
	//dp[1][1] = 1
	//dp[1][2] = 1
	//dp[1][f] = 1
	//dp[2][f] = dp[1][f - 1]  + dp[1][f - 2]  + dp[1][1] ....
	//dp[i][j] = use i dice and get number j
	for i := 1;i <= f;i++{
		dp[1][i] = 1
	}
	for i := 2;i <= d;i++{
		for j := i;j <= i * f && j <= target;j++{
			for k := 1;k <= f && j - k > 0;k++{
				dp[i][j] += dp[i - 1][j - k]
				dp[i][j] %= 1000000007
			}
		}
	}
	return dp[d][target]
}