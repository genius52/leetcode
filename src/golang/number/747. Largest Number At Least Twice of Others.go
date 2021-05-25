package number

func dominantIndex(nums []int) int {
	var l int = len(nums)
	var first int = 0
	var second int = 0
	var index int = -1
	for i := 0;i < l;i++{
		if nums[i] > first{
			second = first
			first = nums[i]
			index = i
		}else if nums[i] > second{
			second = nums[i]
		}
	}
	if first >= second * 2{
		return index
	}
	return -1
}