package array

//DP from bottom to top
func rob198(nums []int) int {
	l := len(nums)
	if l == 0{
		return 0
	}
	var rob_last_last int = 0
	var rob_last int = nums[0]
	for i := 1;i < l;i++{
		tmp := max_int(rob_last_last + nums[i],rob_last)
		rob_last_last = rob_last
		rob_last = tmp
	}
	return rob_last
}