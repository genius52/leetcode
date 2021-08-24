package tree

import "math"

func find_min(arr []int)int{
	var min_val int = 2147483647
	var min_idx int = -1
	for i,v := range arr{
		if v < min_val{
			min_val = v
			min_idx = i
		}
	}
	return min_idx
}

func mctFromLeafValues2(arr []int) int{
	var res int = 0
	for len(arr) > 1{
		idx := find_min(arr)
		if idx == 0{
			res += arr[idx] * arr[idx + 1]
		}else if idx == (len(arr) - 1){
			res += arr[idx] * arr[idx - 1]
		}else{
			res += arr[idx] * min_int(arr[idx - 1],arr[idx + 1])
		}
		arr = append(arr[:idx], arr[idx+1:]...)
	}
	return res
}

func dp_arr(arr []int,max_records [][]int,memo [][]int,left,right int)int{
	if left >= right{
		return 0
	}
	if memo[left][right] != 0{
		return memo[left][right]
	}
	var res int = 1<<30
	for i := left;i < right;i++{
		res = int(math.Min(float64(res),float64(max_records[left][i] * max_records[i+1][right] + dp_arr(arr,max_records,memo,left,i) + dp_arr(arr,max_records,memo,i+1,right))))
	}
	memo[left][right] = res
	return res
}

func mctFromLeafValues(arr []int) int {
	var memo [][]int = make([][]int,len(arr))
	var max_records [][]int = make([][]int,len(arr))
	for i := 0;i < len(arr);i++{
		max_records[i] = make([]int,len(arr))
		memo[i] = make([]int,len(arr))
		max_records[i][i] = arr[i]
		for j := i+1;j < len(arr);j++{
			max_records[i][j] = int(math.Max(float64(arr[j]),float64(max_records[i][j-1])))
		}
	}
	return dp_arr(arr,max_records,memo,0,len(arr)-1)
}
