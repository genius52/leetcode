package number

func maxOperations(nums []int) int {
	var l int = len(nums)
	var cnt int = 1
	for i := 2; i < l-1; i += 2 {
		if nums[i]+nums[i+1] == nums[i-2]+nums[i-1] {
			cnt++
		} else {
			break
		}
	}
	return cnt
}
