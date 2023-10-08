package array

func maximumTripletValue(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				val := int64(nums[i]-nums[j]) * int64(nums[k])
				if val > res {
					res = val
				}
			}
		}
	}
	return res
}
