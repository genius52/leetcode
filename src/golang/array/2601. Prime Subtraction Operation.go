package array

import "sort"

func PrimeSubOperation(nums []int) bool {
	var l int = len(nums)
	var primes [1001]bool //primes[i] = true 表示合数
	for i := 2; i <= 1000; i++ {
		if !primes[i] {
			for j := i * 2; j <= 1000; j += i {
				primes[j] = true
			}
		}
	}
	var prime_nums []int
	for i := 1000; i >= 2; i-- {
		if !primes[i] {
			prime_nums = append(prime_nums, i)
		}
	}
	var l2 int = len(prime_nums)
	var pre_num int = 0
	for i := 0; i < l; i++ {
		if nums[i] <= pre_num {
			return false
		}
		idx := sort.Search(l2, func(j int) bool {
			return nums[i] > pre_num+prime_nums[j]
		})
		if idx == l2 {
			pre_num = nums[i]
		} else {
			pre_num = nums[i] - prime_nums[idx]
		}
	}
	return true
}
