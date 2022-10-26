package array

func calc_minCost(nums []int,cost []int,target int)int64{
	var res int64 = 0
	for i := 0;i < len(nums);i++{
		res += int64(abs_int(nums[i] - target) * cost[i])
	}
	return res
}

func minCost(nums []int, cost []int) int64 {
	var l int = len(nums)
	var low int = 1 << 31 - 1
	var high int = 0
	for i := 0;i < l;i++{
		if nums[i] < low{
			low = nums[i]
		}
		if nums[i] > high{
			high = nums[i]
		}
	}
	if low == high{
		return 0
	}
	var res int64 = 1 << 63 - 1
	for low < high{
		mid1 := (low + high)/2
		mid2 := mid1 + 1
		ret1 := calc_minCost(nums,cost,mid1)
		ret2 := calc_minCost(nums,cost,mid2)
		res = min_int64_number(res,ret1,ret2)
		if ret1 < ret2{
			high = mid1
		}else{
			low =  mid2
		}
	}
	return res
}