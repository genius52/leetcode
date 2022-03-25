package array

func merge88(nums1 []int, m int, nums2 []int, n int)  {
	var idx1 int = m - 1
	var idx2 int = n - 1
	var right int = m + n - 1
	for right >= 0{
		if idx1 < 0{
			nums1[right] = nums2[idx2]
			idx2--
		}else if idx2 < 0{
			nums1[right] = nums1[idx1]
			idx1--
		}else if nums1[idx1] > nums2[idx2]{
			nums1[right] = nums1[idx1]
			idx1--
		}else{
			nums1[right] = nums2[idx2]
			idx2--
		}
		right--
	}
}