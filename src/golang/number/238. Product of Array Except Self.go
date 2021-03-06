package number

func ProductExceptSelf(nums []int) []int {
	var l int = len(nums)
	var left []int = make([]int,l)
	var right []int = make([]int,l)
	left[0] = nums[0]
	right[l - 1] = nums[l - 1]
	for i := 1;i < l;i++{
		left[i] = left[i - 1] * nums[i]
		right[l - 1 - i] = right[l - i] * nums[l - 1 - i]
	}
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		if i == 0{
			res[i] = right[i + 1]
		}else if i == l - 1{
			res[i] = left[i - 1]
		}else{
			res[i] = left[i - 1] * right[i + 1]
		}
	}
	return res
}