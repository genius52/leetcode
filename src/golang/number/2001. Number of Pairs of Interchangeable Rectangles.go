package number

func interchangeableRectangles(rectangles [][]int) int64 {
	var l int = len(rectangles)
	var record map[float64]int64 = make(map[float64]int64)
	var res int64 = 0
	for i := 0;i < l;i++{
		y := rectangles[i][1]
		x := rectangles[i][0]
		key := float64(y)/float64(x)
		if cnt,ok := record[key];ok{
			res += cnt
			record[key]++
		}else{
			record[key] = 1
		}
	}
	return res
}