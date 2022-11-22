package number

func unequalTriplets(nums []int) int {
	var res int = 0
	var l int = len(nums)
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if nums[i] == nums[j]{
				continue
			}
			for k := j + 1;k < l;k++{
				if nums[k] != nums[j] && nums[k] != nums[i]{
					res++
				}
			}
		}
	}
	return res
}