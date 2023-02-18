package array

import "sort"

func MinCost2561(basket1 []int, basket2 []int) int64 {
	var l int = len(basket1)
	var total map[int]int = make(map[int]int)
	var cnt1 map[int]int = make(map[int]int)
	var min_num int = 1 << 31 - 1
	for i := 0;i < l;i++{
		cnt1[basket1[i]]++
		total[basket1[i]]++
		total[basket2[i]]++
		if basket1[i] < min_num{
			min_num = basket1[i]
		}
		if basket2[i] < min_num{
			min_num = basket2[i]
		}
	}
	var diff []int
	for k,v1 := range total{
		if v1 | 1 == v1{
			return -1
		}
		target := v1/2
		if v2,ok := cnt1[k];ok{
			if v2 != target{
				for i := 0;i < abs_int(v2 - target);i++{
					diff = append(diff,k)
				}
			}
		}else{
			for i := 0;i < target;i++{
				diff = append(diff,k)
			}
		}
	}
	sort.Slice(diff, func(i, j int) bool {
		return diff[i] < diff[j]
	})
	var res int64 = 0
	right := len(diff)
	for i := 0;i < right;i++{
		if min_num * 2 <= diff[i]{
			res += int64(min_num * 2)
		}else{
			res += int64(diff[i])
		}
		right--
	}
	return res
}