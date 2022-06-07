package array

func minMaxGame(nums []int) int {
	for len(nums) != 1{
		var l int = len(nums)
		var nums2 []int = make([]int,l/2)
		for i := 0;i < l/2;i++{
			if (i | 1) == i{//odd
				nums2[i] = max_int(nums[i * 2],nums[i * 2 + 1])
			}else{
				nums2[i] = min_int(nums[i * 2],nums[i * 2 + 1])
			}
		}
		nums = nums2
	}
	return nums[0]
}