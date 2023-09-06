package array

//func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
//	var l int = len(nums)
//	var prefix []int = make([]int, l)
//	var cnt int = 0
//	for i := 0; i < l; i++ {
//		if (nums[i] % modulo) == k {
//			cnt++
//		}
//		prefix[i] = cnt % k
//	}
//	var res int64 = 0
//
//	return res
//}

func CountInterestingSubarrays(nums []int, modulo int, k int) int64 {
	var l int = len(nums)
	for i := 0; i < l; i++ {
		if nums[i]%modulo == k {
			nums[i] = 1
		}
	}
	var mod_cnt map[int]int = make(map[int]int)
	var one_cnt int = 0
	var res int64 = 0
	for i := 0; i < l; i++ {
		one_cnt += nums[i]
		m := one_cnt % modulo
		if m == k {
			res++
		}
		target := k - m
		if target < 0 {
			target += modulo
		}
		res += int64(mod_cnt[target])
		mod_cnt[m]++
	}
	return res
}
