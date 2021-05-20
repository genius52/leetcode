package number

func PivotIndex(nums []int) int {
	var l int = len(nums)
	var sum int = 0
	for i := 0;i < l;i++{
		sum += nums[i]
	}
	var cur_sum int = 0
	var index int = 0
	for index < l{
		if cur_sum == (sum - nums[index] - cur_sum){
			return index
		}
		cur_sum += nums[index]
		index++
	}
	return -1
}