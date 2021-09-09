package diagram

func abs_int(n int)int{
	if n < 0{
		return -n
	}
	return n
}

func minTimeToVisitAllPoints(points [][]int) int {
	var res int = 0
	for pos := 1; pos < len(points);pos++{
		res += max_int(abs_int(points[pos][0] - points[pos-1][0]),abs_int(points[pos][1] - points[pos-1][1]))
	}
	return res
}