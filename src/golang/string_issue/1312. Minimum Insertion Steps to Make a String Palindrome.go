package string_issue

//Input: s = "mbadm"
//Output: 2
//Explanation: String can be "mbdadbm" or "mdbabdm".
func MinInsertions1312(s string) int {
	var l int = len(s)
	var dp [][]int = make([][]int,l)//dp[i][j]:从i到j的最长回文序列长度
	for i := 0;i < l;i++{
		dp[i] = make([]int,l)
		dp[i][i] = 1
	}
	var max_len int = 1
	for i := 1;i < l;i++{
		for j := i - 1;j >= 0;j--{
			if s[i] == s[j]{
				if j == i - 1{
					dp[j][i] = 2
				}else{
					dp[j][i] = dp[j + 1][i - 1] + 2
				}
			}else{
				dp[j][i] = max_int(dp[j + 1][i],dp[j][i - 1])
			}
			max_len = max_int(max_len,dp[j][i])
		}
	}
	return l - max_len
}