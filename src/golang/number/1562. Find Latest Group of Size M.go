package number

func FindLatestStep(arr []int, m int) int {
	var l int = len(arr)
	var to_left map[int]int = make(map[int]int)//start: length
	var to_right map[int]int = make(map[int]int)//end: length
	var len_cnt []int = make([]int,l + 1)
	var steps int = -1
	for i := 0;i < l;i++{
		last := arr[i] - 1
		next := arr[i] + 1
		if pre_len,ok1 := to_left[last];ok1{
			len_cnt[pre_len]--
			if next_len,ok2 := to_right[next];ok2{ //前后都是1
				len_cnt[next_len]--
				new_start := arr[i] - pre_len
				new_end := arr[i] + next_len
				new_len := pre_len + next_len + 1
				to_left[new_end] = new_len
				to_right[new_start] = new_len
				len_cnt[new_len]++
				delete(to_left,last)
				delete(to_right,next)
			}else{//前面是1
				new_len := pre_len + 1
				to_left[arr[i]] = new_len
				new_start := arr[i] - pre_len
				to_right[new_start] = new_len
				len_cnt[new_len]++
				delete(to_left,last)
			}
		}else{
			if next_len,ok2 := to_right[next];ok2{//后面是1
				len_cnt[next_len]--
				new_len := next_len + 1
				end := arr[i] + next_len
				to_left[end] = new_len
				to_right[arr[i]] = new_len
				len_cnt[new_len]++
				delete(to_right,next)
			}else{//前后都不是1
				to_left[arr[i]] = 1
				to_right[arr[i]] = 1
				len_cnt[1]++
			}
		}
		if len_cnt[m] > 0{
			steps = i + 1
		}
	}
	return steps
}