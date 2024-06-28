package array

func minOperations3191(nums []int) int {
	var l int = len(nums)
	var res int = 0
	for i := 0; i < l-2; i++ {
		if nums[i] == 0 {
			nums[i+1] = (nums[i+1] + 1) % 2
			nums[i+2] = (nums[i+2] + 1) % 2
			res++
		}
	}
	if nums[l-1] != 1 || nums[l-2] != 1 {
		return -1
	}
	return res
}
