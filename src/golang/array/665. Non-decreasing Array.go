package array

//nums = [4,2,3]
func CheckPossibility(nums []int) bool {
	var l int = len(nums)
	var index int = 1
	for index < l{
		if nums[index] < nums[index - 1]{
			break
		}
		index++
	}
	if index == l{
		return true
	}
	if index > 1{
		if nums[index] < nums[index - 2]{//change nums[index] to a big number
			nums[index] = nums[index - 1]
		}
	}
	index++
	for index < l{
		if nums[index] < nums[index - 1]{
			return false
		}
		index++
	}
	return true
}