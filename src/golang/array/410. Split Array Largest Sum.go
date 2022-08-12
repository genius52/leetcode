package array

func check_splitArray(nums []int,l int,m int,target int)bool{
	var groups int = 1
	var sum int = 0
	for i := 0;i < l;i++{
		if nums[i] > target{
			return false
		}
		sum += nums[i]
		if sum > target{
			groups++
			sum = nums[i]
			if groups > m{
				return false
			}
		}
	}
	return true
}

func SplitArray(nums []int, m int) int {
	var l int = len(nums)
	var low int = 1
	var high int = 0
	for i := 0;i < l;i++{
		if nums[i] > low{
			low = nums[i]
		}
		high += nums[i]
	}
	if m == 1{
		return high
	}
	var res int = low
	for low <= high{
		mid := (low + high)/2 //假设和的最大值为mid
		ret := check_splitArray(nums,l,m,mid)
		if ret{
			high = mid - 1
			res = mid
		}else{
			low = mid + 1
			res = mid + 1
		}
	}
	return res
}