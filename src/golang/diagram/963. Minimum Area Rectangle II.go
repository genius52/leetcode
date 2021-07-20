package diagram

import (
	"math"
	"strconv"
)

func distance(x1 int,y1 int,x2 int,y2 int)int{
	return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2)
}

func MinAreaFreeRect(points [][]int) float64 {
	var l int = len(points)
	if l < 4{
		return 0
	}
	var record map[string][]int = make(map[string][]int)
	//把中心点相同的
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			dis := distance(points[i][0],points[i][1],points[j][0],points[j][1])
			s1 := strconv.FormatFloat((float64(points[i][0]) + float64(points[j][0]))/2,'f',6,64)
			s2 := strconv.FormatFloat((float64(points[i][1]) + float64(points[j][1]))/2,'f',6,64)
			key := s1 + "_" + s2 + "_" + strconv.Itoa(dis)
			record[key] = append(record[key],i)
			record[key] = append(record[key],j)
		}
	}
	var res float64 = 2147483476
	for _,index := range record{
		if len(index) <= 2{
			continue
		}
		var cur_l int = len(index)
		for i := 0;i < cur_l;i += 2{
			for j := i + 2;j < cur_l;j += 2{
				p1 := points[index[i]]
				p2 := points[index[i + 1]]
				p3 := points[index[j]]
				//p4 := points[index[j + 1]]
				dis1 := math.Sqrt(float64(distance(p1[0],p1[1],p3[0],p3[1])))
				dis2 := math.Sqrt(float64(distance(p2[0],p2[1],p3[0],p3[1])))
				square := dis1 * dis2
				if square < res && square != 0{
					res = square
				}
			}
		}
	}
	if res == 2147483476{
		return 0
	}
	return res
}