package array

func FindIndices2905(nums []int, indexDifference int, valueDifference int) []int {
	var l int = len(nums)
	var min_val int = 1<<31 - 1
	var min_val_idx int = 0
	var max_val int = -1 << 31
	var max_val_idx int = 0
	var left int = 0
	var right int = indexDifference
	for right < l {
		if nums[left] < min_val {
			min_val = nums[left]
			min_val_idx = left
		}
		if nums[left] > max_val {
			max_val = nums[left]
			max_val_idx = left
		}
		if abs_int(nums[right]-min_val) >= valueDifference {
			return []int{min_val_idx, right}
		}
		if abs_int(nums[right]-max_val) >= valueDifference {
			return []int{max_val_idx, right}
		}
		left++
		right++
	}
	return []int{-1, -1}
}
