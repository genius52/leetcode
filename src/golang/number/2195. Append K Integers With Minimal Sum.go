package number

import "sort"

//输入：nums = [5,6], k = 6
//输出：25
//解释：在该解法中，向数组中追加的两个互不相同且未出现的正整数是 1 、2 、3 、4 、7 和 8 。
//nums 最终元素和为 5 + 6 + 1 + 2 + 3 + 4 + 7 + 8 = 36 ，这是所有情况中的最小值。
//所以追加到数组中的两个整数之和是 1 + 2 + 3 + 4 + 7 + 8 = 25 ，所以返回 25 。
func MinimalKSum(nums []int, k int) int64 {
	nums = append(nums,0)
	nums = append(nums,1 << 31 - 1)
	sort.Ints(nums)
	var sum int64 = 0
	var l int = len(nums)
	for i := 1;i < l;i++{
		space_cnt := nums[i] - nums[i - 1] - 1
		if space_cnt < 1{
			continue
		}
		if k <= space_cnt{
			cur := int64(k * (nums[i - 1] + 1) + k * (k - 1)/2)
			sum += cur
			break
		}else{
			cur := int64(space_cnt * (nums[i - 1] + 1) + space_cnt * (space_cnt - 1)/2)
			k -= space_cnt
			sum += cur
		}
	}
	return sum
}