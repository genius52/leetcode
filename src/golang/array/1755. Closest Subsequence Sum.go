package array

import "sort"

//Input: nums = [7,-9,15,-2], goal = -5
//Output: 1
//Explanation: Choose the subsequence [7,-9,-2], with a sum of -4.
//The absolute difference is abs(-4 - (-5)) = abs(1) = 1, which is the minimum.
func make_minAbsDifference(nums []int) []int {
	var l int = len(nums)
	sum := make([]int, 1<<l) //sum[i]:
	for i := 1; i < 1<<l; i++ {
		for j := 0; j < l; j++ {
			if i|(1<<j) == i {
				sum[i] = sum[i-1<<j] + nums[j]
			}
		}
	}
	return sum
}

func MinAbsDifference(nums []int, goal int) int {
	var l int = len(nums)
	sum1 := make_minAbsDifference(nums[:l/2])
	sum2 := make_minAbsDifference(nums[l/2:])
	sort.Ints(sum1)
	sort.Slice(sum2, func(i, j int) bool {
		return sum2[i] > sum2[j]
	})
	var l1 int = len(sum1)
	var l2 int = len(sum2)
	var res int = abs_int(goal)
	for i := 0; i < l1; i++ {
		res = min_int(res, abs_int(sum1[i]-goal))
	}
	for i := 0; i < l2; i++ {
		res = min_int(res, abs_int(sum2[i]-goal))
	}
	var a int = 0
	var b int = 0
	for a < l1 && b < l2 {
		cur := sum1[a] + sum2[b]
		res = min_int(res, abs_int(cur-goal))
		if cur < goal {
			a++
		} else if cur > goal {
			b++
		} else {
			break
		}
	}
	return res
}
