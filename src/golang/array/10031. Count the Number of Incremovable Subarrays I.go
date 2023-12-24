package array

func IncremovableSubarrayCount(nums []int) int {
	var l int = len(nums)
	var res int = 1
	for sub_len := 1; sub_len <= l-1; sub_len++ {
		for left := 0; left < l && left+sub_len <= l; left++ { //remove nums[left:left + sub_len]
			var increase bool = true
			for i := 0; i < l; i++ {
				if i >= left && i < left+sub_len {
					continue
				}
				if i == (left + sub_len) {
					last := left - 1
					if last >= 0 {
						if nums[i] <= nums[last] {
							increase = false
							break
						} else {
							continue
						}
					} else {
						continue
					}
				}
				if i > 0 && nums[i] <= nums[i-1] {
					increase = false
					break
				}
			}
			if increase {
				res++
			}
		}
	}
	return res
}
