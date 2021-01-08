package string_issue

func NumberOfSubstrings(s string) int {
	var l int = len(s)
	var dp [][3]int = make([][3]int,l)//dp[i][0] = 从i位置开始出现的第一个a
	for i := 0;i <= 2;i++{
		dp[l - 1][i] = -1
	}
	dp[l - 1][s[l - 1] - 'a'] = l - 1
	for i := l - 2;i >= 0;i--{
		for j := 0;j <= 2;j++{
			dp[i][j] = dp[i + 1][j]
		}
		dp[i][s[i] - 'a'] = i
	}
	var res int = 0
	var begin int = 0
	var end int = 2
	for begin < l - 2 &&  end < l{
		if dp[begin][0] == -1 || dp[begin][1] == -1 || dp[begin][2] == -1{
			break
		}
		if dp[begin][0] > end || dp[begin][1] > end || dp[begin][2] > end{
			end++
		}else{
			res += l - end
			begin++
		}
	}
	return res
}