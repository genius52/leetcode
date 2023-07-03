package number

func isFascinating(n int) bool {
	var check [10]int
	var n2 int = n * 2
	var n3 int = n * 3
	for _, c := range []int{n, n2, n3} {
		for c > 0 {
			m := c % 10
			if m == 0 {
				return false
			}
			check[m]++
			if check[m] > 1 {
				return false
			}
			c /= 10
		}
	}
	for i := 1; i <= 9; i++ {
		if check[i] != 1 {
			return false
		}
	}
	return true
}
