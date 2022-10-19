package array

func MinimizeArrayValue(nums []int) int {
	var l int = len(nums)
	var less_cnt int = 0
	var left_max int = nums[0]
	for i := 1;i < l;i++{
		if nums[i] > left_max{
			increase := nums[i] - left_max
			if increase > less_cnt{
				increase -= less_cnt
				nums_cnt := i + 1
				left_max += increase / nums_cnt
				if increase % nums_cnt > 0{
					left_max++
					less_cnt = nums_cnt - increase % nums_cnt
				}else{
					less_cnt = 0
				}
			}else{
				less_cnt -= increase
			}
		}else{
			less_cnt += left_max - nums[i]
		}
	}
	return left_max
}