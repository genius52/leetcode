package array

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	//var l int = len(nums1)
	//var prefix1 []int = make([]int,l + 1)
	//var prefix2 []int = make([]int,l + 1)
	//var sum1 int = 0
	//var sum2 int = 0
	//for i := 0;i < l;i++{
	//	sum1 += nums1[i]
	//	sum2 += nums2[i]
	//	prefix1[i + 1] = prefix1[i] + nums1[i]
	//	prefix2[i + 1] = prefix2[i] + nums2[i]
	//}
	//var res int = max_int(sum1,sum2)
	//for right := 1;right < l;right++{
	//	for left := right - 1;left >= 0;left--{
	//
	//	}
	//}
	//return res

	var l int = len(nums1)
	var sum1 int = 0
	var sum2 int = 0
	for i := 0;i < l;i++{
		sum1 += nums1[i]
		sum2 += nums2[i]
	}
	return 0
}