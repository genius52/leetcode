package array

func getCommon(nums1 []int, nums2 []int) int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var i int = 0
	var j int = 0
	for i < l1 && j < l2{
		if nums1[i] == nums2[j]{
			return nums1[i]
		}else if nums1[i] < nums2[j]{
			i++
		}else if nums1[i] > nums2[j]{
			j++
		}
	}
	return -1
}