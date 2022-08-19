package array

import "sort"

func mergesort_reversePairs(nums []int,left int,right int)int{
	if left == right{
		return 0
	}
	mid := (left + right)/2
	cnt := mergesort_reversePairs(nums,left,mid) + mergesort_reversePairs(nums,mid + 1,right)
	var idx1 int = left
	var idx2 int = mid + 1
	for idx1 <= mid{
		for idx2 <= right && nums[idx1] > nums[idx2] * 2{
			idx2++
		}
		cnt += idx2 - 1 - mid
		idx1++
	}
	sort.Ints(nums[left:right + 1])
	return cnt
}

func reversePairs(nums []int) int {
	var l int = len(nums)
	return mergesort_reversePairs(nums,0,l - 1)
}