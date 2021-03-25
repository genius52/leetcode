package array

func numberOfArithmeticSlices2(A []int) int{
	var l int = len(A)
	if l < 3{
		return 0
	}
	var dp []int = make([]int,l)//dp[i] = from 0 to i,the count of arithmetic slices
	if A[1] - A[0] == A[2] - A[1]{
		dp[2] = 1
	}
	for i := 3;i < l;i++{
		if A[i] - A[i - 1] == A[i - 1] - A[i - 2]{
			dp[i] = dp[i - 1] + 1
		}
	}
	var total int = 0
	for i := 2;i < l;i++{
		total += dp[i]
	}
	return total
}