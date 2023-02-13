package array

//如果一个四元组 (i, j, k, l) 满足以下条件，我们称它是上升的：
//
//0 <= i < j < k < l < n 且
//nums[i] < nums[k] < nums[j] < nums[l] 。

//输入：nums = [1,3,2,4,5]
//输出：2
//解释：
//- 当 i = 0 ，j = 1 ，k = 2 且 l = 3 时，有 nums[i] < nums[k] < nums[j] < nums[l] 。
//- 当 i = 0 ，j = 1 ，k = 2 且 l = 4 时，有 nums[i] < nums[k] < nums[j] < nums[l] 。
func countQuadruplets2552(nums []int) int64 {
	var l2 int = len(nums)
	var dp_j []int64 = make([]int64,l2)
	var res int64 = 0
	for l := 0;l < l2;l++{
		for j := 0;j < l;j++{
			if nums[l] > nums[j]{
				res += dp_j[j]
			}
		}
		var small_than_k int64 = 0
		for j := 0;j < l;j++{
			if nums[j] > nums[l]{
				dp_j[j] += small_than_k
			}
			if nums[j] < nums[l] {
				small_than_k++
			}
		}
	}
	return res
}