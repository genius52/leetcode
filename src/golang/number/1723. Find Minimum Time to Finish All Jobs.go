package number

//输入：jobs = [1,2,4,7,8], k = 2
//输出：11
//解释：按下述方式分配工作：
//1 号工人：1、2、8（工作时间 = 1 + 2 + 8 = 11）
//2 号工人：4、7（工作时间 = 4 + 7 = 11）
//最大工作时间是 11 。
func MinimumTimeRequired(jobs []int, k int) int {
	var l int = len(jobs)
	total_jobs := 1 << l
	var values []int = make([]int, total_jobs) //事先计算好所有Job组合的值
	for i := 1; i < total_jobs; i++ {
		for j := 0; j < l; j++ {
			if ((1 << j) & i) != 0 {
				values[i] += jobs[j]
			}
		}
	}
	//dp[i][j] 表示：前 i 个工人为了完成作业子集 j，需要花费的最大工作时间的最小值
	var dp [][]int = make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, 1<<l)
		for j := 0; j < total_jobs; j++ {
			dp[i][j] = values[j]
		}
	}
	for i := 1; i < k; i++ {
		for state := 1; state <= 1<<l-1; state++ {
			// 枚举 j 的全部子集
			min_val := 2147483647
			//sub和j 互补
			for sub_state := state; sub_state != 0; sub_state = (sub_state - 1) & state {
				cur := max_int(dp[i-1][state-sub_state], values[sub_state])
				min_val = min_int(min_val, cur)
			}
			dp[i][state] = min_val
		}
	}
	return dp[k-1][1<<l-1]
}
