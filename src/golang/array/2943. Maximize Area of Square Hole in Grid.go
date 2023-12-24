package array

import "sort"

func MaximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)
	var max_h int = 1
	var max_v int = 1
	var cur_h int = 1
	var cur_v int = 1
	var l_h int = len(hBars)
	var l_v int = len(vBars)
	for i := 0; i < len(hBars); i++ {
		if i == 0 {
			cur_h++
		} else {
			if hBars[i] == hBars[i-1]+1 {
				cur_h++
			} else {
				cur_h = 2
			}
		}
		if cur_h > max_h {
			max_h = cur_h
		}
	}
	if l_h == 1 {
		max_h = 2
	}
	for i := 0; i < len(vBars); i++ {
		if i == 0 {
			cur_v++
		} else {
			if vBars[i] == vBars[i-1]+1 {
				cur_v++
			} else {
				cur_v = 2
			}
		}
		if cur_v > max_v {
			max_v = cur_v
		}
	}
	if l_v == 1 {
		max_v = 2
	}
	edge := min_int(max_h, max_v)
	return edge * edge
}
