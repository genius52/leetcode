package array

func minimumRightShifts(nums []int) int {
	var l int = len(nums)
	var pre_max int = nums[0]
	var decrease_cnt int = 0
	var decrease_idx int = -1
	for i := 1; i < l; i++ {
		if decrease_cnt > 0 && nums[i] > pre_max {
			return -1
		}
		if nums[i] < nums[i-1] {
			decrease_cnt++
			if decrease_cnt > 1 {
				return -1
			}
			decrease_idx = i
		}
	}
	if decrease_idx == -1 {
		return 0
	}
	return l - decrease_idx
}
