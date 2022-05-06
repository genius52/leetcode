package array

func wiggleMaxLength(nums []int) int {
	var l int = len(nums)
	if l <= 1{
		return l
	}
	var pre_increase_len int = 1
	var pre_decrease_len int = 1
	for i := 1;i < l;i++{
		if nums[i] > nums[i - 1]{
			pre_increase_len = pre_decrease_len + 1
		}else if nums[i] < nums[i - 1]{
			pre_decrease_len = pre_increase_len + 1
		}
	}
	if pre_increase_len > pre_decrease_len{
		return pre_increase_len
	}else{
		return pre_decrease_len
	}
}