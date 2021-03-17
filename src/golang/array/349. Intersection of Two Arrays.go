package array

func intersection(nums1 []int, nums2 []int) []int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var record1 [1001]bool
	for i := 0;i < l1;i++{
		record1[nums1[i]] = true
	}
	var record2 [1001]bool
	for i := 0;i < l2;i++{
		if record1[nums2[i]]{
			record2[nums2[i]] = true
		}
	}
	var res []int
	for i := 0;i <= 1000;i++{
		if record2[i]{
			res = append(res,i)
		}
	}
	return res
}