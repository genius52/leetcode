package array

import "math"

func maximumSubarraySum(nums []int, k int) int64 {
	var l int = len(nums)
	var sum int64 = 0
	var res int64 = math.MinInt64
	var val_presum map[int][]int64 = make(map[int][]int64)
	for i := 0; i < l; i++ {
		sum += int64(nums[i])
		target1 := nums[i] - k
		if presums, ok := val_presum[target1]; ok {
			for _, pre_sum := range presums {
				cur := sum - pre_sum
				if cur > res {
					res = cur
				}
			}
		}
		target2 := nums[i] + k
		if presums, ok := val_presum[target2]; ok {
			for _, pre_sum := range presums {
				cur := sum - pre_sum
				if cur > res {
					res = cur
				}
			}
		}
		val_presum[nums[i]] = append(val_presum[nums[i]], sum-int64(nums[i]))
	}
	if res == math.MinInt64 {
		res = 0
	}
	return res
}
