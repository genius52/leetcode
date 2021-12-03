package number

func countGoodRectangles(rectangles [][]int) int {
	var l int = len(rectangles)
	var max_len int = 0
	var max_cnt int = 0
	for i := 0;i < l;i++{
		var cur int = min_int(rectangles[i][0],rectangles[i][1])
		if cur > max_len{
			max_len = cur
			max_cnt = 1
		}else if cur == max_len{
			max_cnt++
		}
	}
	return max_cnt
}