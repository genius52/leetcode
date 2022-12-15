package array

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	var l int = len(nums1)
	var same_cnt map[int]int = make(map[int]int)
	var max_same_cnt int = 0
	for i := 0;i < l;i++{
		if nums1[i] == nums2[i]{
			same_cnt[nums1[i]]++
			if same_cnt[nums1[i]] > max_same_cnt{
				max_same_cnt = same_cnt[nums1[i]]
			}
		}
	}
	if max_same_cnt > l/2{
		return -1
	}
	
}