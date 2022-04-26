package diagram

import "sort"

func CountRectangles(rectangles [][]int, points [][]int) []int {
	var l int = len(points)
	var record [101][]int
	for i := 0;i < len(rectangles);i++{
		record[rectangles[i][1]] = append(record[rectangles[i][1]],rectangles[i][0])
	}
	for i := 1;i <= 100;i++{
		sort.Ints(record[i])
	}
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		x := points[i][0]
		y := points[i][1]
		var cnt int = 0
		for j := y;j <= 100;j++{
			sub_len := len(record[j])
			if sub_len == 0 || x > record[j][sub_len - 1]{
				continue
			}
			left := 0
			right := sub_len - 1
			for left < right{
				mid := (left + right)/2
				if record[j][mid] < x{
					left = mid + 1
				}else{
					right = mid
				}
			}
			cnt += len(record[j]) - left
		}
		res[i] = cnt
	}
	return res
}