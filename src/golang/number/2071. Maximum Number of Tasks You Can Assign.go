package number

import "sort"

func check_maxTaskAssign(tasks []int,workers []int,pills int, strength int)bool{
	var l int = len(tasks)
	var w2 []int = make([]int,l)
	copy(w2,workers)
	for i := l - 1;i >= 0;i--{
		var work_len int = len(w2)
		if (pills == 0 && w2[work_len - 1] < tasks[i]) || (w2[work_len - 1] + strength < tasks[i]){
			return false
		}
		if w2[work_len - 1] >= tasks[i]{
			w2 = w2[:work_len - 1]
			continue
		}
		var target int = tasks[i]
		if pills > 0{
			pills--
			target = tasks[i] - strength
		}
		var left int = 0
		var right int = work_len - 1
		//找到第一个大于等于tasks[i]
		var res int = 0
		for left < right{
			mid := (left + right)/2
			if workers[mid] < target{
				left = mid + 1
				res = mid + 1
			}else if workers[mid] == target{
				res = mid
				break
			}else{
				right = mid
				res = mid
			}
		}
		w2 = append(w2[:res],w2[res+1:]...)
	}
	return true
}

func MaxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	var res int = 0
	var l1 int = len(tasks)
	var l2 int = len(workers)
	var low int = 0
	var high int = l1
	for low <= high{
		mid := (low + high)/2
		var start int = 0
		if l2 - mid > 0{
			start = l2 - mid
		}
		ret := check_maxTaskAssign(tasks[:mid],workers[start:],pills,strength)
		if ret{
			low = mid + 1
			res = mid
		}else{
			high = mid - 1
			res = mid - 1
		}
	}
	return res
}