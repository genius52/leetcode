package number

import "sort"

func minimumAverage(nums []int) float64 {
	var l int = len(nums)
	sort.Ints(nums)
	var res float64 = 1<<63 - 1
	for i := 0; i < l/2; i++ {
		average := float64(nums[i]+nums[l-1-i]) / 2
		if average < res {
			res = average
		}
	}
	return res
}
