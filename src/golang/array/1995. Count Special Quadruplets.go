package array

func countQuadruplets(nums []int) int {
	var l int = len(nums)
	var res int = 0
	for i := 3;i < l;i++{
		for third := i - 1;third >= 2;third--{
			for second := third - 1;second >= 1;second--{
				for first := second - 1;first >= 0;first--{
					if nums[first] + nums[second] + nums[third] == nums[i]{
						res++
					}
				}
			}
		}
	}
	return res
}