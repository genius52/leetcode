package array

import "sort"

func specialArray(nums []int) int {
	var l int = len(nums)
	for n := l;n >= 0;n--{
		var cnt int = 0
		for i := 0;i < l;i++{
			if nums[i] >= n{
				cnt++
				if cnt > n{
					break
				}
			}
		}
		if cnt == n{
			return n
		}
	}
	return -1
}


func specialArray2(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	for i := 1;i <= l;i++{
		if i != l{
			if nums[l - i] >= i && nums[l - i - 1] < i{
				return i
			}
		}else{
			if nums[l - i] >= i{
				return i
			}
		}
	}
	return -1
}