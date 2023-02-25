package array

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var idx1 int = 0
	var idx2 int = 0
	var res [][]int
	for idx1 < l1 || idx2 < l2 {
		if idx1 == l1 {
			res = append(res, nums2[idx2])
			idx2++
		} else if idx2 == l2 {
			res = append(res, nums1[idx1])
			idx1++
		} else {
			if nums1[idx1][0] == nums2[idx2][0] {
				res = append(res, []int{nums1[idx1][0], nums1[idx1][1] + nums2[idx2][1]})
				idx1++
				idx2++
			} else if nums1[idx1][0] < nums2[idx2][0] {
				res = append(res, nums1[idx1])
				idx1++
			} else if nums1[idx1][0] > nums2[idx2][0] {
				res = append(res, nums2[idx2])
				idx2++
			}
		}
	}
	return res
}
