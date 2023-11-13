package number

func distributeCandies2928(n int, limit int) int {
	var res int = 0
	for i := 0; i <= limit && i <= n; i++ {
		for j := 0; j <= limit && j <= n-i; j++ {
			for k := 0; k <= limit && k <= n-i-j; k++ {
				if (i + j + k) == n {
					res++
				}
			}
		}
	}
	return res
}
