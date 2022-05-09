package array

func isMonotonic(nums []int) bool {
	var l int = len(nums)
	if l <= 2{
		return true
	}
	var idx int = 1
	for idx < l && nums[idx] == nums[idx - 1]{
		idx++
	}
	if idx == l{
		return true
	}
	var increase bool = nums[idx] > nums[idx - 1]
	idx++
	for idx < l{
		if increase{
			if nums[idx] < nums[idx - 1]{
				return false
			}
		}else{
			if nums[idx] > nums[idx - 1]{
				return false
			}
		}
		idx++
	}
	return true
}

func isMonotonic2(nums []int) bool{
	var l int = len(nums)
	if l <= 2{
		return true
	}
	var increase bool = true
	var decrease bool = true
	for i := 1;i < l;i++{
		if nums[i - 1] < nums[i]{
			decrease = false
		}
		if nums[i - 1] > nums[i]{
			increase = false
		}
		if !increase && !decrease{
			return false
		}
	}
	return true
}