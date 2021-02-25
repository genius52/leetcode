package diagram

import "strconv"

func gcd(num1 int,num2 int)int{
	if num1 != 0{
		return gcd(num2 % num1, num2)
	}else{
		return num2
	}
}

func MaxPoints(points [][]int) int {
	var l int = len(points)
	if l < 3{
		return 2
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
			a /= g
			b /= g
			//c /= c
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