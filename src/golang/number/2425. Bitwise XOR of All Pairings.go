package number

func xorAllNums(nums1 []int, nums2 []int) int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var cnt [32]int
	for i := 0;i < l1;i++{
		var offset int = 0
		for nums1[i] > 0{
			if nums1[i] | 1 == nums1[i]{
				cnt[offset] += l2
			}
			offset++
			nums1[i] >>= 1
		}
	}
	for i := 0;i < l2;i++{
		var offset int = 0
		for nums2[i] > 0{
			if nums2[i] | 1 == nums2[i]{
				cnt[offset] += l1
			}
			offset++
			nums2[i] >>= 1
		}
	}
	var res int = 0
	for i := 0;i < 32;i++{
		if cnt[i] | 1 == cnt[i]{
			res |= 1 << i
		}
	}
	return res
}