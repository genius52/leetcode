package array

import "sort"

func findLonely(nums []int) []int {
	var l int = len(nums)
	if l == 1{
		return []int{nums[0]}
	}
	sort.Ints(nums)
	var res []int
	for i := 0;i < l;i++{
		if i == 0{
			if nums[i] != nums[i + 1] && nums[i + 1] != nums[i] + 1{
				res = append(res,nums[i])
			}
		}else if i == l - 1{
			if nums[i] != nums[i - 1] && nums[i - 1] != nums[i] - 1{
				res = append(res,nums[i])
			}
		}else {
			if nums[i] != nums[i - 1] && nums[i] != nums[i + 1] && nums[i - 1] != nums[i] - 1 && nums[i + 1] != nums[i] + 1{
				res = append(res,nums[i])
			}
		}
	}
	return res
}