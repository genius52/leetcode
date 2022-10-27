package array

import (
	"sort"
	"strconv"
)

func dfs_makeArrayIncreasing(arr1 []int,l1 int,idx1 int,arr2 []int,l2 int,idx2 int,pre int,memo map[string]int)int{
	if idx1 == l1{
		return 0
	}
	key := strconv.Itoa(idx1) + "," + strconv.Itoa(idx2) + "," + strconv.Itoa(pre)
	if _,ok := memo[key];ok{
		return memo[key]
	}
	if arr1[idx1] <= pre{ //替换
		for idx2 < l2 && arr2[idx2] <= pre{
			idx2++
		}
		if idx2 == l2{
			memo[key] = -1
			return -1
		}
		ret := dfs_makeArrayIncreasing(arr1,l1,idx1 + 1,arr2,l2,idx2 + 1,arr2[idx2],memo)
		if ret == -1{
			return ret
		}
		return ret + 1
	}else{
		//replace current
		var ret1 int = -1
		for idx2 < l2 && arr2[idx2] <= pre{
			idx2++
		}
		if idx2 != l2{
			ret1 = dfs_makeArrayIncreasing(arr1,l1,idx1 + 1,arr2,l2,idx2 + 1,arr2[idx2],memo)
			if ret1 != -1{
				ret1++
			}
		}
		//No action
		ret2 := dfs_makeArrayIncreasing(arr1,l1,idx1 + 1,arr2,l2,idx2,arr1[idx1],memo)
		if ret1 == -1{
			memo[key] = ret2
		}else if ret2 == -1{
			memo[key] = ret1
		}else{
			memo[key] = min_int(ret1,ret2)
		}
		return memo[key]
	}
}

func MakeArrayIncreasing(arr1 []int, arr2 []int) int {
	var l1 int = len(arr1)
	var l2 int = len(arr2)
	var memo map[string]int = make(map[string]int)
	sort.Ints(arr2)
	return dfs_makeArrayIncreasing(arr1,l1,0,arr2,l2,0,-1,memo)
}