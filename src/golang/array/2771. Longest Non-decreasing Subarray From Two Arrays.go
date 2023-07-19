package array

func MaxNonDecreasingLength(nums1 []int, nums2 []int) int {
	var l int = len(nums1)
	var dp1 []int = make([]int, l)
	var dp2 []int = make([]int, l)
	dp1[0] = 1
	dp2[0] = 1
	var res int = 1
	for i := 1; i < l; i++ {
		dp1[i] = 1
		if nums1[i] >= nums1[i-1] {
			dp1[i] = max_int(dp1[i], 1+dp1[i-1])
		}
		if nums1[i] >= nums2[i-1] {
			dp1[i] = max_int(dp1[i], 1+dp2[i-1])
		}
		dp2[i] = 1
		if nums2[i] >= nums1[i-1] {
			dp2[i] = max_int(dp2[i], 1+dp1[i-1])
		}
		if nums2[i] >= nums2[i-1] {
			dp2[i] = max_int(dp2[i], 1+dp2[i-1])
		}
		res = max_int(res, dp1[i])
		res = max_int(res, dp2[i])
	}
	return res
}
