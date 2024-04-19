package diagram

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var l int = len(points)
	var cnt int = 1
	var pre_x int = points[0][0]
	for i := 1; i < l; i++ {
		if points[i][0]-pre_x > w {
			cnt++
			pre_x = points[i][0]
		}
	}
	return cnt
}
