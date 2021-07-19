package string_issue

//Input: strs = ["babca","bbazb"]
//Output: 3
//Explanation: After deleting columns 0, 1, and 4, the final array is strs = ["bc", "az"].
func MinDeletionSize3(strs []string) int {
	var rows int = len(strs)
	var columns int = len(strs[0])
	var dp []int = make([]int,columns) //dp[i] 以i结尾，最大升序长度
	for i := 0;i < columns;i++{
		dp[i] = 1
	}
	var res int = 0
	for i := 0;i < columns;i++{
		for j := i - 1;j >= 0;j--{
			var increase bool = true
			for r := 0;r < rows;r++{
				if strs[r][i] < strs[r][j]{
					increase = false
					break
				}
			}
			if increase{
				dp[i] = max_int(dp[i],1 + dp[j])
			}
		}
		res = max_int(res,dp[i])
	}
	return columns - res
}