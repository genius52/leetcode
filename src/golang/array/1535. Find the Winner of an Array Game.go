package array

import (
	"math"
)

func getWinner(arr []int, k int) int {
	var l int = len(arr)
	var biggest int = arr[0]
	var win_cnt int = 0
	for i := 1;i < l;i++{
		if(arr[i] > biggest){
			win_cnt = 1
			biggest = arr[i]
		}else{
			win_cnt++
		}
		if(win_cnt == k){
			return biggest
		}
	}
	return biggest
}

func GetWinner(arr []int, k int) int {
	var win_cnt int = 0
	var l int = len(arr)
	if(k >= l){
		var res int = math.MinInt32
		for i := 0;i < l;i++{
			if(arr[i] > res){
				res = arr[i]
			}
		}
		return res
	}
	for win_cnt < k{
		if(arr[0] > arr[1]){
			win_cnt++
			if(win_cnt == k){
				return arr[0]
			}
			loser := arr[1]
			arr = append(arr[:1], arr[2:]...)
			arr = append(arr[:], loser)
			//arr = newSlice
		}else{
			win_cnt = 1
			if(k == 1){
				return arr[1]
			}
			loser := arr[0]
			arr = append(arr[1:],loser)
			//arr = newSlice
		}
	}
	return -1
}


func GetWinner2(arr []int, k int) int {
	var l int = len(arr)
	if(k >= l){
		var res int = math.MinInt32
		for i := 0;i < l;i++{
			if(arr[i] > res){
				res = arr[i]
			}
		}
		return res
	}
	var biggest int = arr[0]
	var win_cnt int = 0
	for i := 1;i < l;i++{
		if(arr[i] > biggest){
			win_cnt = 1
			biggest = arr[i]
		}else{
			win_cnt++
		}
		if(win_cnt == k){
			return biggest
		}
	}
	return biggest
}