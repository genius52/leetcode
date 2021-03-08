package diagram

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	var l int = len(points)
	var index int = -1
	var distance int = 2147483647
	for i := 0;i < l;i++{
		if x != points[i][0] && y != points[i][1]{
			continue
		}
		cur_dis := int(math.Abs(float64(x) - float64(points[i][0]))) + int(math.Abs(float64(y) - float64(points[i][1])))
		if cur_dis < distance{
			distance = cur_dis
			index = i
		}
	}
	return index
}