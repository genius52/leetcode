package number

import "sort"

//Input: n = 7, cuts = [1,3,4,5]
//Output: 16
//Explanation: Using cuts order = [1, 3, 4, 5] as in the input leads to the following scenario:
func dp_minCost(cuts []int,l int,left int,right int,memo [][]int)int{
	if right - left <= 1{
		return 0
	}
	if memo[left][right] != 0{
		return memo[left][right]
	}
	var res int = 2147483647
	for i := left + 1;i < right;i++{
		cur := cuts[right] - cuts[left] + dp_minCost(cuts,l,left,i,memo) + dp_minCost(cuts,l,i,right,memo)
		if cur < res{
			res = cur
		}
	}
	memo[left][right] = res
	return res
}

//DP from top to bottom
func MinCost(n int, cuts []int) int{
	cuts = append(cuts,0)
	cuts = append(cuts,n)
	sort.Ints(cuts)
	var l int = len(cuts)
	var memo [][]int = make([][]int,l)//dp[left][right]:从i到j切割的最小值
	for i := 0;i < l;i++{
		memo[i] = make([]int,l)
	}
	return dp_minCost(cuts,l,0,l - 1,memo)
}

func MinCost2(n int, cuts []int) int {
	cuts = append(cuts,0)
	cuts = append(cuts,n)
	sort.Ints(cuts)
	var l int = len(cuts)
	var dp [][]int = make([][]int,l)//dp[left][right]:从i到j切割的最小值
	for i := 0;i < l;i++{
		dp[i] = make([]int,l)
	}
	for cur_len := 2;cur_len < l;cur_len++{
		for left := 0;left < l;left++{
			right := left + cur_len
			if right >= l{
				break
			}
			dp[left][right] = 2147483647 //从cuts[left] 拼成cuts[right]的cost
			for i := left + 1;i < right;i++{
				dp[left][right] = min_int(dp[left][right],cuts[right] - cuts[left] + dp[left][i] + dp[i][right])
			}
		}
	}
	return dp[0][l - 1]
}