package number

import "math"

func DistinctPrimeFactors(nums []int) int {
	var record map[int]bool = make(map[int]bool)
	var prime_num []int
	for i := 2;i <= 1000;i++{
		var is_prime bool = true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++{
			if i % j == 0{
				is_prime = false
				break
			}
		}
		if is_prime{
			prime_num = append(prime_num,i)
		}
	}
	for _,n := range nums{
		for j := 0;j < len(prime_num);j++{
			for n >= prime_num[j] && (n % prime_num[j]) == 0{
				n /= prime_num[j]
				record[prime_num[j]] = true
			}
			if n < prime_num[j]{
				break
			}
		}
	}
	return len(record)
}