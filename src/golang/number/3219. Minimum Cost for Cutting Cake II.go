package number

import "sort"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	sort.Ints(horizontalCut)
	sort.Ints(verticalCut)
	var idx_h int = m - 2
	var idx_v int = n - 2
	var vertical_pre_cnt int = 1   //垂直方向要切几次
	var horizontal_pre_cnt int = 1 //水平方向要切几次
	var res int64 = 0
	for idx_h >= 0 && idx_v >= 0 {
		if horizontalCut[idx_h] >= verticalCut[idx_v] {
			res += int64(horizontalCut[idx_h] * horizontal_pre_cnt)
			vertical_pre_cnt++
			idx_h--
		} else {
			res += int64(verticalCut[idx_v] * vertical_pre_cnt)
			horizontal_pre_cnt++
			idx_v--
		}
	}
	for idx_h >= 0 {
		res += int64(horizontalCut[idx_h] * horizontal_pre_cnt)
		idx_h--
		vertical_pre_cnt++
	}
	for idx_v >= 0 {
		res += int64(verticalCut[idx_v] * vertical_pre_cnt)
		idx_v--
		horizontal_pre_cnt++
	}
	return res
}
