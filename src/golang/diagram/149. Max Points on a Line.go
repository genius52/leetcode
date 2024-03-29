package diagram

import "strconv"

func gcd(a int,b int)int{
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func MaxPoints(points [][]int) int {
	var l int = len(points)
	if l <= 2{
		return l
	}
	var record map[string]int = make(map[string]int)
	var max_cnt int = 0
	for i := 0;i < l;i++{
		for j := 0;j < l;j++{
			if j == i{
				continue
			}
			//ax + by + c = 0
			a := points[j][1] - points[i][1]
			b := points[i][0] - points[j][0]
			c := points[j][0] * points[i][1] - points[i][0] * points[j][1]
			//a, b, c = y2 - y1, x1 - x2, x2 * y1 - x1 * y2
			g := gcd(a,b)
			g2 := gcd(g,c)
			a /= g2
			b /= g2
			c /= g2
			key := strconv.Itoa(a) + "," + strconv.Itoa(b) + "," + strconv.Itoa(c)
			if _,ok := record[key];ok{
				continue
			}else{
				record[key] = 2
			}
			for k := 0;k < l;k++{
				if k == i || k == j{
					continue
				}
				if (points[k][0] * a + points[k][1] * b + c) == 0{
					record[key]++
				}
			}
			if record[key] > max_cnt{
				max_cnt = record[key]
			}
		}
	}
	return max_cnt
}