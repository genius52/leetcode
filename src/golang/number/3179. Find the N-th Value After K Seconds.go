package number

func ValueAfterKSeconds(n int, k int) int {
	var MOD int = 1e9 + 7
	var prefix_sum []int = make([]int, n)
	for i := 0; i < n; i++ {
		prefix_sum[i] = 1
	}
	for k > 0 {
		for i := 1; i < n; i++ {
			prefix_sum[i] += prefix_sum[i-1]
			prefix_sum[i] %= MOD
		}
		k--
	}
	return prefix_sum[n-1]
}
