package diagram

import "math"

//Input: radius = 1, x_center = 0, y_center = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
//Output: true
//Explanation: Circle and rectangle share the point (1,0)
func CheckOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	x1 -= x_center
	x2 -= x_center
	y1 -= y_center
	y2 -= y_center
	var x_min int = 0
	if (x1 >= 0 && x2 <= 0) || (x1 <= 0 && x2 >= 0){
		x_min = 0
	}else{
		x_min = int(math.Min(math.Abs(float64(x1)),math.Abs(float64(x2))))
	}
	var y_min int = 0
	if (y1 >= 0 && y2 <= 0) || (y1 <= 0 && y2 >= 0) {
		y_min = 0
	}else{
		y_min = int(math.Min(math.Abs(float64(y1)),math.Abs(float64(y2))))
	}

	return x_min * x_min + y_min * y_min <= radius * radius
}