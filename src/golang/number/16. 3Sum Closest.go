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
	var begin int = 0
	var end int = l - 1
	for begin < end - 1{
		mid := begin + 1
		var end2 int = end
		for mid < end2{
			sum := nums[begin] + nums[mid] + nums[end2]
			diff := int(math.Abs(float64(sum - target)))
			if diff < min_gap{
				min_gap = diff
				res = sum
			}
			if sum > target{
				end2--
			}else{
				mid++
			}
		}
		begin++
	}
	return res
}