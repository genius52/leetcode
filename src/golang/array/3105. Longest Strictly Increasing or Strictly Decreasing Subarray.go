package array

func longestMonotonicSubarray(nums []int) int {
	var l int = len(nums)
	var res int = 1
	var increase int = 1
	var decrease int = 1
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			increase++
			decrease = 1
			if increase > res {
				res = increase
			}
		} else if nums[i] < nums[i-1] {
			decrease++
			increase = 1
			if decrease > res {
				res = decrease
			}
		} else {
			increase = 1
			decrease = 1
		}
	}
	return res
}
