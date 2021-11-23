package array

import "sort"

func check_finish(tasks [][]int,l int,total int)bool{
	for i := 0;i < l;i++{
		if total < tasks[i][1]{
			return false
		}
		total -= tasks[i][0]
	}
	return true
}

func MinimumEffort(tasks [][]int) int {
	var low int = 0
	var high int = 0
	var l int = len(tasks)
	for i := 0;i < l;i++{
		low += tasks[i][0]
		high += tasks[i][1]
	}
	sort.Slice(tasks, func(i, j int) bool {
		diff1 := tasks[i][1] - tasks[i][0]
		diff2 := tasks[j][1] - tasks[j][0]
		if diff1 == diff2{
			return tasks[i][0] > tasks[j][0]
		}else{
			return diff1 > diff2
		}
	})
	for low < high{
		mid := (low + high)/2
		ret := check_finish(tasks,l,mid)
		if ret{
			high = mid
		}else{
			low = mid + 1
		}
	}
	return low
}