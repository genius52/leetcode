package array

func buildArray(nums []int) []int {
	var l int = len(nums)
	var res []int = make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = nums[nums[i]]
	}
	return res
}
