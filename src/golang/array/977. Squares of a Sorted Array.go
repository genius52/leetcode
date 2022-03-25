package array

func sortedSquares(nums []int) []int {
	var l int = len(nums)
	var res []int = make([]int,l)
	var low int = 0
	var high int = l - 1
	var index int = l - 1
	for low <= high{
		if nums[low] >= 0{
			res[index] = nums[high] * nums[high]
			high--
		}else{
			if nums[high] > -nums[low]{
				res[index] = nums[high] * nums[high]
				high--
			}else{
				res[index] = nums[low] * nums[low]
				low++
			}
		}
		index--
	}
	return res
}