package number

import "math"

// 质数的平方是特殊数
func nonSpecialCount(l int, r int) int {
	var limit int = int(math.Sqrt(float64(r)))
	var is_prime []bool = make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		is_prime[i] = true
	}
	//is_prime[2] = false
	for i := 2; i <= limit; i++ {
		if is_prime[i] {
			for j := i * i; j <= limit; j += i {
				is_prime[j] = false
			}
		}
	}
	var not_special_cnt int = 0
	for i := 2; i <= limit; i++ {
		if is_prime[i] {
			n := i * i
			if n >= l && n <= r {
				not_special_cnt++
			}
		}
	}
	return r - l + 1 - not_special_cnt
}
