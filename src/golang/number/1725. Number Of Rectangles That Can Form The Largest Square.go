package number

func countGoodRectangles(rectangles [][]int) int {
	var l int = len(rectangles)
	var max_len int = 0
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		var cur int = min_int(rectangles[i][0],rectangles[i][1])
		if cur > max_len{
			max_len = cur
		}
		if _,ok := record[cur];ok{
			record[cur]++
		}else{
			record[cur] = 1
		}
	}
	return record[max_len]
}