package array

import "sort"

func MinAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	var l int = len(nums1)
	var sum int64 = 0
	var MOD int = 1e9 + 7
	var sorted []int = make([]int,l)
	copy(sorted,nums1)
	sort.Ints(sorted)
	var max_decrease int = 0
	for i := 0;i < l;i++{
		if nums1[i] == nums2[i]{
			continue
		}
		diff1 := abs_int(nums1[i] - nums2[i])
		sum += int64(diff1)
		var left int = 0
		var right int = l - 1
		for left < right{
			mid := (left + right)/2
			if sorted[mid] == nums2[i]{
				left = mid
				break
			}else if sorted[mid] < nums2[i]{
				left = mid + 1
			}else if sorted[mid] > nums2[i]{
				right = mid
			}
		}
		diff2 := abs_int(sorted[left] - nums2[i])
		if left - 1 >= 0{
			diff2 = min_int(diff2,abs_int(sorted[left - 1] - nums2[i]))
		}
		if left + 1 < l{
			diff2 = min_int(diff2,abs_int(sorted[left + 1] - nums2[i]))
		}
		max_decrease = max_int(max_decrease,diff1 - diff2)
	}
	return int(sum - int64(max_decrease)) % MOD
}