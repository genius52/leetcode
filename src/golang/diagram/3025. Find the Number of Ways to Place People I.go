package diagram

import "sort"

func NumberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		} else {
			return points[i][1] >= points[j][1]
		}
	})
	var l int = len(points)
	var cnt int = 0
	for i := 0; i < l; i++ {
		var find bool = false
		var close_y int = -1
		for j := i + 1; j < l; j++ {
			if points[i][1] >= points[j][1] {
				if !find {
					find = true
					close_y = points[j][1]
					cnt++
				} else {
					if points[j][1] > close_y {
						cnt++
					}
					close_y = max_int(close_y, points[j][1])
				}
			}
		}
	}
	return cnt
}
