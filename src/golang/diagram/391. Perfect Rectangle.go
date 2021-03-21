package diagram

import "math"

//rectangles = [
//  [1,1,3,3],
//  [3,1,4,2],
//  [3,2,4,4],
//  [1,3,2,4],
//  [2,3,3,4]
//]

//static uint8_t is_rect_intersect(int x01, int x02, int y01, int y02,
//	int x11, int x12, int y11, int y12)
//{
//	int zx = abs(x01 + x02 -x11 - x12);
//	int x  = abs(x01 - x02) + abs(x11 - x12);
//	int zy = abs(y01 + y02 - y11 - y12);
//	int y  = abs(y01 - y02) + abs(y11 - y12);
//	if(zx <= x && zy <= y)
//		return 1;
//	else
//		return 0;
//}
func is_cross(rect1 []int,rect2 []int)bool{
	var zx int = int(math.Abs(float64(rect1[0] + rect1[2] - rect2[0] - rect2[2])))
	var x int = int(math.Abs(float64(rect1[0] - rect1[2]))) + int(math.Abs(float64(rect2[0] - rect2[2])))
	var zy int = int(math.Abs(float64(rect1[1] + rect1[3] - rect2[1] - rect2[3])))
	var y int = int(math.Abs(float64(rect1[1] - rect1[3]))) + int(math.Abs(float64(rect2[1] - rect2[3])))
	if zx < x && zy < y{
		return true
	}else{
		return false
	}
}

func IsRectangleCover(rectangles [][]int) bool {
	var l int = len(rectangles)
	var left int = 2147483647
	var right int = -2147483648
	var bottom int = 2147483647
	var top int = -2147483648
	var total_area int = 0
	for i := 0;i < l;i++{
		area := (rectangles[i][2] - rectangles[i][0]) * (rectangles[i][3] - rectangles[i][1])
		total_area += area
		if rectangles[i][0] < left{
			left = rectangles[i][0]
		}
		if rectangles[i][2] > right{
			right = rectangles[i][2]
		}
		if rectangles[i][1] < bottom{
			bottom = rectangles[i][1]
		}
		if rectangles[i][3] > top{
			top = rectangles[i][3]
		}
	}
	if (right - left) * (top - bottom) != total_area{
		return false
	}
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if is_cross(rectangles[i],rectangles[j]){
				return false
			}
		}
	}
	return true
}