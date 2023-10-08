package array

func MaxSubarrays(nums []int) int {
	var l int = len(nums)
	var zero_cnt [32]int
	var groups int = 1
	
	for i := 0; i < l; i++ {
		for j := 0; j < 32; j++ {
			if nums[i]|1 != nums[i] {
				zero_cnt[j]++
			}
		}

	}
	return groups
}
