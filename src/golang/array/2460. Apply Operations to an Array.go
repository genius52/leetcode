package array

func applyOperations(nums []int) []int {
	var l int = len(nums)
	for i := 0;i < l - 1;{
		if nums[i] != nums[i + 1]{
			i++
			continue
		}
		nums[i] *= 2
		nums[i + 1] = 0
		i += 2
	}
	var left int = 0
	for left < l {
		for left < l && nums[left] != 0{
			left++
		}
		var right int = left + 1
		for right < l && nums[right] == 0{
			right++
		}
		if left == l || right == l{
			break
		}
		nums[left] = nums[right]
		nums[right] = 0
		left++
	}
	return nums
}