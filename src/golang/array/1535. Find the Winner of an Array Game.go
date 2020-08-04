package array

import (
	"math"
)

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