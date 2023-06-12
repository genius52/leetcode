package array

func semiOrderedPermutation(nums []int) int {
	var l int = len(nums)
	var one_idx int = -1
	var n_idx int = -1
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			one_idx = i
			if n_idx != -1 {
				break
			}
		}
		if nums[i] == l {
			n_idx = i
			if one_idx != -1 {
				break
			}
		}
	}
	if one_idx > n_idx {
		return one_idx + l - 1 - n_idx - 1
	}
	return one_idx + l - 1 - n_idx
}
