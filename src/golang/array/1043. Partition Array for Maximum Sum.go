package array

//DP from bottom to top
func maxSumAfterPartitioning2(A []int, K int) int{
	var l int = len(A)
	var dp []int = make([]int,l + 1)//dp[i] = max sum from A[i] to A[i - 1]
	for i := 1;i <= l;i++{
		var max_val int = 0
		for j := 1;j <= K && (i - j) >= 0;j++{
			max_val = max_int(max_val,A[i - j])//A[i - k]..A[i - 1]的最大值
			dp[i] = max_int(dp[i],max_val * j + dp[i - j])
		}
	}
	return dp[l]
}

//DP from top to bottom
func max_int_slice(nums []int) int{
	var max int = 0
	for _,val := range nums{
		if val > max{
			max = val
		}
	}
	return max
}

func max_sum(A []int,begin int,K int,records map[int]int) int{
	var max int = 0
	if begin >= len(A) {
		return max
	}
	for i := 1;i <= K;i++{
		if (begin + i)  > len(A) {
			break
		}
		var sum int = 0
		if val,ok := records[begin+i];ok{
			sum = val
		}else{
			sum = max_sum(A,begin + i,K,records)
			records[begin + i] = sum
		}
		cur_sum := max_int_slice(A[begin:begin+i]) * (i) + sum
		if cur_sum > max{
			max = cur_sum
		}
	}
	return max
}

func maxSumAfterPartitioning(A []int, K int) int {
	var records map[int]int = make(map[int]int,len(A))
	return max_sum(A,0,K,records)
}