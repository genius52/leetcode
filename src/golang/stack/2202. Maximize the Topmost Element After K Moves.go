package stack

func maximumTop(nums []int, k int) int {
	if k == 0{
		return nums[0]
	}
	var l int = len(nums)
	if l <= 1{
		if (k | 1) == k{
			return -1
		}else{
			return nums[0]
		}
	}
	if k == 1{
		return nums[1]
	}
	if k > l{
		var res int = 0
		for i := 0;i < l;i++{
			if nums[i] > res{
				res = nums[i]
			}
		}
		return res
	}else if k < l{
		var res int = nums[k]
		for i := 0;i < k - 1;i++{
			if nums[i] > res{
				res = nums[i]
			}
		}
		return res
	}else{
		var res int = 0
		for i := 0;i < l - 1;i++{
			if nums[i] > res{
				res = nums[i]
			}
		}
		return res
	}
}