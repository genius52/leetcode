package array

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	var l int = len(nums1)
	var sum1 int = 0
	var sum2 int = 0
	for i := 0;i < l;i++{
		sum1 += nums1[i]
		sum2 += nums2[i]
	}
	var max_increase1 int = 0
	var max_increase2 int = 0
	var pre_inc1 int = 0
	var pre_inc2 int = 0
	for i := 0;i < l;i++{
		inc1 := nums2[i] - nums1[i]
		inc2 := nums1[i] - nums2[i]
		pre_inc1 = max_int(pre_inc1 + inc1,inc1)
		pre_inc2 = max_int(pre_inc2 + inc2,inc2)
		max_increase1 = max_int(max_increase1,pre_inc1)
		max_increase2 = max_int(max_increase2,pre_inc2)
	}
	return max_int(sum1 + max_increase1,sum2 + max_increase2)
}