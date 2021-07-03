package array

import (
	"math"
	"sort"
)

//Input: nums = [2,1,3]
//Output: 6
//Explanation:
//Subsequences are [1], [2], [3], [2,1], [2,3], [1,3], [2,1,3].
//The corresponding widths are 0, 0, 0, 1, 1, 2, 2.
//The sum of these widths is 6.
func SumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int64 = 0
	for i := 0;i < l;i++{
		//nums[i] as the biggest number
		res = (res + int64(nums[i]) * (int64(math.Pow(2,float64(i))) % 1000000007) - int64(nums[i]) * (int64(math.Pow(2,float64(l - i - 1))) % 1000000007)) % 1000000007
	}
	return int(res + 1000000007) % 1000000007
}