package number

import (
	"sort"
)

//arr = [60864,25176,27249,21296,20204], target = 56803
func FindBestValue(arr []int, target int) int{
	sort.Ints(arr)
	var l int = len(arr)
	if (arr[0] * l) >= target{
		res1 := target / l
		res2 := target / l + 1
		if (target - res1 * l) <= (res2 * l - target){
			return res1
		} else{
			return res2
		}
	}
	if (arr[l - 1] * l) <= target{
		return arr[l - 1]
	}
	var prefix int = 0
	var cur_sum int = 0
	for i := 0;i < l;i++{
		cur_sum = prefix + arr[i] * (l - i)
		if cur_sum >= target{
			res1 := (target - prefix)/ (l - i)
			res2 := res1 + 1
			val1 := target - prefix - res1 * (l - i)
			val2 := prefix + res2 * (l - i) - target
			if val1 <= val2{
				return res1
			}else{
				return res2
			}
		}
		prefix += arr[i]
	}
	return arr[l - 1]
}

func findBestValue(arr []int, target int) int{
	var l int = len(arr)
	var left int = 0
	var right int = 0
	var sum int = 0
	for i := 0;i < l;i++{
		if arr[i] > right{
			right = arr[i]
		}
		sum += arr[i]
	}
	var min_diff int = 2147483647
	var min_diff_val int = right
	for left <= right{
		mid := (left + right)/2
		var cur_sum int = 0
		for i := 0;i < l;i++{
			if arr[i] > mid{
				cur_sum += mid
			}else{
				cur_sum += arr[i]
			}
		}
		if cur_sum < target{
			left = mid + 1
		}else{
			right = mid - 1
		}
		cur_diff := abs_int(cur_sum - target)
		if cur_diff < min_diff{
			min_diff_val = mid
			min_diff = cur_diff
		}else if cur_diff == min_diff{
			if mid < min_diff_val{
				min_diff_val = mid
			}
		}
	}
	return min_diff_val
}