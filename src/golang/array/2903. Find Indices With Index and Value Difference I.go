package array

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	var l int = len(nums)
	var res []int
	for i := 0; i+indexDifference < l; i++ {
		for j := i + indexDifference; j < l; j++ {
			if abs_int(nums[i]-nums[j]) >= valueDifference {
				res = []int{i, j}
				return res
			}
		}
	}
	return []int{-1, -1}
}

func findIndices2(nums []int, indexDifference int, valueDifference int) []int {
	var l int = len(nums)
	var res []int
	var min_idx int = 0
	var max_idx int = 0
	for i := indexDifference; i < l; i++ {
		left := i - indexDifference
		if nums[left] < nums[min_idx] {
			min_idx = left
		}
		if nums[left] > nums[max_idx] {
			max_idx = left
		}
		if abs_int(nums[i]-nums[min_idx]) >= valueDifference {
			res = []int{min_idx, i}
			return res
		}
		if abs_int(nums[i]-nums[max_idx]) >= valueDifference {
			res = []int{max_idx, i}
			return res
		}
	}
	return []int{-1, -1}
}
