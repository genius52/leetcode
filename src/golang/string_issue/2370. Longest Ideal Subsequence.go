package string_issue

//LIS
func LongestIdealString(s string, k int) int {
	var l int = len(s)
	var dp [26]int //dp[i]
	dp[s[0] - 'a'] = 1
	for i := 1;i < l;i++{
		var cur_max int = 1
		for j := 0;j < 26;j++{
			diff := abs_int(int(s[i] - 'a') - j)
			if diff <= k{
				cur_max = max_int(cur_max,dp[j] + 1)
			}
		}
		dp[s[i] - 'a'] = cur_max
	}
	var res int = 0
	for i := 0;i < 26;i++{
		if dp[i] > res{
			res = dp[i]
		}
	}
	return res
}

//func LongestIdealString(s string, k int) int {
//	var l int = len(s)
//	var dp []int = make([]int,l)
//	dp[0] = 1
//	var res int = 1
//	for i := 1;i < l;i++{
//		var cur_max int = 1
//		for j := i - 1;j >= 0;j--{
//			diff := abs_int(int(s[i]) - int(s[j]))
//			if diff <= k{
//				cur_max = max_int(cur_max,dp[j] + 1)
//			}
//		}
//		dp[i] = cur_max
//		if cur_max > res{
//			res = cur_max
//		}
//	}
//	return res
//}