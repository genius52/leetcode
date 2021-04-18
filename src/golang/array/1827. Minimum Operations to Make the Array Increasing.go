package array

func minOperations827(nums []int) int {
	var l int = len(nums)
	var steps int = 0
	var pre int = nums[0]
	for i := 1;i < l;i++{
		if nums[i] <= pre{
			steps += pre - nums[i] + 1
			pre = pre + 1
		}else{
			pre = nums[i]
		}
	}
	return steps
}