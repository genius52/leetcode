package number

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var l3 int = len(nums3)
	var record1 [101]bool
	for i := 0;i < l1;i++{
		record1[nums1[i]] = true
	}
	var record2 [101]bool
	for i := 0;i < l2;i++{
		record2[nums2[i]] = true
	}
	var record3 [101]bool
	for i := 0;i < l3;i++{
		record3[nums3[i]] = true
	}
	var res []int
	for i := 1;i <= 100;i++{
		if (record1[i] && record2[i]) || (record1[i] && record3[i]) ||
			(record2[i] && record3[i]){
			res = append(res,i)
		}
	}
	return res
}