package array

import "sort"

func check_minimumIncompatibility(data []int,l int)int{
	sort.Ints(data)
	for i := 1;i < l;i++{
		if data[i] == data[i - 1]{
			return -1
		}
	}
	return data[l - 1] - data[0]
}

func dp_minimumIncompatibility(nums []int,l int,k int,target_len int,cur []int,status int,memo map[int]map[int]int)int{
	if k == 0{
		return 0
	}
	if _,ok := memo[k];!ok{
		memo[k] = make(map[int]int)
	}else{
		if _,ok2 := memo[k][status];ok2{
			return memo[k][status]
		}
	}
	var res int = 2147483647
	if len(cur) == target_len{
		diff := check_minimumIncompatibility(cur,target_len)
		if diff == -1{
			res = -1
		}else{
			next := dp_minimumIncompatibility(nums,l,k - 1,target_len,[]int{},status,memo)
			if next == -1{
				res = -1
			}else{
				res = diff + next
			}
		}
	}else{
		for i := 0;i < l;i++{
			if nums[i] != 0{
				back := nums[i]
				cur = append(cur,nums[i])
				nums[i] = 0
				ret := dp_minimumIncompatibility(nums,l,k,target_len,cur,status | (1 << i),memo)
				cur = cur[:len(cur) - 1]
				if ret != -1{
					res = min_int(res,ret)
				}
				nums[i] = back
			}
		}
	}
	if res == 2147483647{
		res = -1
	}
	memo[k][status] = res
	return res
}

func MinimumIncompatibility(nums []int, k int) int {
	var l int = len(nums)
	var status int = 0
	var memo map[int]map[int]int = make(map[int]map[int]int)
	return dp_minimumIncompatibility(nums,l,k,l / k,[]int{},status,memo)
}