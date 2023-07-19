package array

func CheckArray(nums []int, k int) bool {
	var l int = len(nums)
	var pre_decrease int = 0
	for i := 0; i < l; i++ {
		if pre_decrease > nums[i] { //over decrease
			return false
		}
		nums[i] -= pre_decrease //1,2,3,4,k = 3
		pre_decrease += nums[i]
		if i >= (k - 1) {
			pre_decrease -= nums[i-k+1]
		}
	}
	return pre_decrease == 0
}
