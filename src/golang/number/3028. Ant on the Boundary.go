package number

func returnToBoundaryCount(nums []int) int {
	var l int = len(nums)
	var cnt int = 0
	var pos int = 0
	for i := 0; i < l; i++ {
		pos += nums[i]
		if pos == 0 {
			cnt++
		}
	}
	return cnt
}
