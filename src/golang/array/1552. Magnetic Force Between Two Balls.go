package array

import (
	"sort"
)

func check_dis(position []int,ball_cnt int,distance int)bool{
	var l int = len(position)
	var pre_num int = position[0]
	var cnt int = 1
	for cur_pos := 1;cur_pos < l;cur_pos++{
		if position[cur_pos] - pre_num >= distance{
			cnt++
			pre_num = position[cur_pos]
		}
	}
	return cnt >= ball_cnt
}

func MaxDistance(position []int, m int) int {
	sort.Ints(position)
	var l int = len(position)
	var high int =  position[l - 1]
	var low int = 0
	var res int = 0
	for low <= high{
		mid := low + (high-low)/2
		ret := check_dis(position,m,mid)
		if ret{
			low = mid + 1
			res = mid
		}else{
			high = mid - 1
		}
	}
	return res
}