package number

import (
	"math"
	"sort"
)

//Input: nums = [-1,2,1,-4], target = 1
//Output: 2
//Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
func ThreeSumClosest(nums []int, target int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var min_gap int = 1<< 31 - 1
	var res int = 0
	var left int = 0
	for left < l - 1{
		mid := left + 1
		var right int = l - 1
		for mid < right{
			sum := nums[left] + nums[mid] + nums[right]
			diff := int(math.Abs(float64(sum - target)))
			if diff < min_gap{
				min_gap = diff
				res = sum
			}
			if sum > target{
				right--
			}else{
				mid++
			}
		}
		left++
	}
	return res
}