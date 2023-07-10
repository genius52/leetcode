package array

func findPrimePairs(n int) [][]int {
	//1,2,3,5,7,11
	var not_prime []bool = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !not_prime[i] {
			for j := 2; i*j <= n; j++ {
				not_prime[i*j] = true
			}
		}
	}
	var res [][]int
	for i := 2; i <= n/2; i++ {
		if !not_prime[i] && !not_prime[n-i] {
			res = append(res, []int{i, n - i})
		}
	}
	return res
}
