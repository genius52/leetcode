package array

import "sort"

func MinOperations2009(nums []int) int {
	var l1 int = len(nums)
	var data map[int]bool = make(map[int]bool)
	for i := 0;i < l1;i++{
		data[nums[i]] = true
	}
	var l2 int = len(data)
	var nums2 []int = make([]int,l2)
	var idx int = 0
	for n,_ := range data{
		nums2[idx] = n
		idx++
	}
	sort.Ints(nums2)
	var left int = 0
	var right int = 0
	var res int = 0
	for left < l2{
		for right < l2 && nums2[right] - nums2[left] <= l1 - 1{
			right++
		}
		res = max_int(res,right - left)
		left++
	}
	return l1  - res
}