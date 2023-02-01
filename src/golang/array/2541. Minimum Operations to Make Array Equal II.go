package array

func minOperations2541(nums1 []int, nums2 []int, k int) int64 {
	var l int = len(nums1)
	var more_cnt int64 = 0
	var less_cnt int64 = 0
	for i := 0;i < l;i++{
		diff := nums1[i] - nums2[i]
		if k == 0{
			if nums1[i] != nums2[i]{
				return -1
			}
			continue
		}
		if diff % k != 0{
			return -1
		}
		if diff > 0{
			more_cnt += int64(diff / k)
		}else if diff < 0{
			less_cnt += int64(-diff / k)
		}
	}
	if more_cnt != less_cnt{
		return -1
	}
	return more_cnt
}