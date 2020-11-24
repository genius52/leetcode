package diagram

import "math"

//Input: [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
//Output: 2
func MinAreaRect(points [][]int) int {
	var xkey_map map[int]map[int]bool = make(map[int]map[int]bool)
	var l int = len(points)
	for i := 0;i < l;i++{
		if _,ok := xkey_map[points[i][0]];ok{
			xkey_map[points[i][0]][points[i][1]] = true
		}else{
			xkey_map[points[i][0]] = make(map[int]bool)
			xkey_map[points[i][0]][points[i][1]] = true
		}
	}
	var res int = math.MaxInt32
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if points[i][0] == points[j][0] || points[i][1] == points[j][1]{
				continue
			}
			if _,ok1 := xkey_map[points[i][0]][points[j][1]];ok1{
				if _,ok2 := xkey_map[points[j][0]][points[i][1]];ok2{
					width := int(math.Abs(float64(points[i][0] - points[j][0])))
					height := int(math.Abs(float64(points[i][1] - points[j][1])))
					area := width * height
					if area < res{
						res = area
					}
				}
			}
		}
	}
	if res == math.MaxInt32{
		return 0
	}
	return res
}