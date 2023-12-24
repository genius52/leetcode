package number

import "sort"

func largestPerimeter2(nums []int) int64 {
	var l int = len(nums)
	sort.Ints(nums)
	var res int64 = -1
	var sum int64 = int64(nums[0] + nums[1])
	for i := 2; i < l; i++ {
		if int64(nums[i]) < sum {
			res = sum + int64(nums[i])
		}
		sum += int64(nums[i])
	}
	return res
}
