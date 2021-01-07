package number

import (
	"math"
	"sort"
)

//arr = [60864,25176,27249,21296,20204], target = 56803
func FindBestValue(arr []int, target int) int {
	var l int = len(arr)
	if l == 1{
		if arr[0] >= target{
			return target
		}else{
			return arr[0]
		}
	}
	sort.Ints(arr)
	var dp []int = make([]int,l)
	var high int = 0
	dp[0] = arr[0]
	for i := 1;i < l;i++{
		dp[i] = arr[i] + dp[i - 1]
		if dp[i] >= target{
			high = arr[i]
		}
	}
	if dp[l - 1] <= target{
		return arr[l - 1]
	}
	if dp[0] >= target{
		return arr[0]
	}
	var diff int = 2147483647
	var pre_num int = arr[l - 1]
	var low int = 0
	for low <= high{
		mid_val := (low + high)/2
		var cur_sum int = 0
		var visit int = 0
		for visit < l{
			if arr[visit] >= mid_val{
				cur_sum += mid_val
			}else{
				cur_sum += arr[visit]
			}
			visit++
		}
		cur_diff := int(math.Abs(float64(cur_sum - target)))
		if cur_diff < diff{
			diff = cur_diff
			pre_num = mid_val
		}else if cur_diff == diff{
			if mid_val < pre_num{
				pre_num = mid_val
			}
		}
		if cur_sum > target{
			high = mid_val - 1
		}else if cur_sum < target{
			low = mid_val + 1
		}else{
			return mid_val
		}
	}
	return pre_num
}