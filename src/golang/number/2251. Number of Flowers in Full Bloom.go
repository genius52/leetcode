package number

import "sort"

func FullBloomFlowers(flowers [][]int, persons []int) []int {
	var start []int
	var end []int
	for _,flower := range flowers{
		start = append(start,flower[0])
		end = append(end,flower[1] + 1)
	}
	sort.Ints(start)
	sort.Ints(end)
	var start_len int = len(start)
	var end_len int = len(end)
	var l int = len(persons)
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		var start_cnt int = 0
		time := persons[i]
		var left int = 0
		var right int = start_len - 1
		if time >= start[right]{
			start_cnt = start_len
		}else if time < start[left]{
			start_cnt = 0
		} else{
			for left < right{//找到第一个大于time的位置
				mid := (left + right)/2
				if start[mid] <= time{
					left = mid + 1
				}else{
					right = mid
				}
			}
			start_cnt = left
		}
		left = 0
		right = end_len - 1
		var end_cnt int = 0
		if time >= end[right]{
			end_cnt = end_len
		}else if time < end[left] {
			end_cnt = 0
		}else{
			var idx int = 0
			for left <= right{//找到第一个小于等于time的位置
				mid := (left + right)/2
				if end[mid] > time{
					right = mid - 1
					idx = mid - 1
				}else{
					left = mid + 1
					idx = mid
				}
			}
			end_cnt = idx + 1
		}
		res[i] = start_cnt - end_cnt
	}
	return res
}