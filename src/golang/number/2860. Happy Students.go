package number

import "sort"

// 这位学生被选中，并且被选中的学生人数 严格大于 nums[i] 。
// 这位学生没有被选中，并且被选中的学生人数 严格小于 nums[i] 。
func countWays2860(nums []int) int {
	var l int = len(nums)
	if l == 1 {
		return 1
	}
	sort.Ints(nums)
	var res int = 0
	if 0 < nums[0] { //一个都不选
		res++
	}
	for i := 0; i < l; i++ {
		//选中nums[0] ... nums[i],
		if i == l-1 {
			if l > nums[i] {
				res++
			}
		} else {
			if (i+1) > nums[i] && (i+1) < nums[i+1] {
				res++
			}
		}
	}
	return res
}
