package number

func IsBoomerang(points [][]int) bool {
	return (points[0][1] - points[2][1]) * (points[0][0] - points[1][0]) != (points[0][1] - points[1][1]) * (points[0][0] - points[2][0])
	//(x1y2+x2y3+x3y1-x1y3-x2y1-x3y2)
	//var res = (points[0][0] * points[1][1] + points[1][0] * points[2][1] +
	//	points[2][0] * points[0][1] - points[0][0] * points[2][1] - points[1][0] * points[0][1] -
	//	points[2][0] * points[1][1])
	//return res != 0
}
