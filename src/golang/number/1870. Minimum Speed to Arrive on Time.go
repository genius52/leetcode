package number

import "math"

//Input: dist = [1,3,2], hour = 2.7
//Output: 3
func MinSpeedOnTime(dist []int, hour float64) int {
	var l int = len(dist)
	if (float64(l) - 1) >= hour{
		return -1
	}
	var sum_distance int = 0
	var high int = 10000000
	for i := 0;i < l;i++{
		sum_distance += dist[i]
		//if high < dist[i]{
		//	high = dist[i]
		//}
	}

	var low int = int(float64(sum_distance)/hour)
	if low == 0{
		low++
	}
	var res int = 0
	for low < high{
		var mid int = low + (high - low) / 2
		var cur_hour float64 = 0
		for i := 0;i < l;i++{
			if i == (l - 1){
				cur_hour += float64(dist[i]) / float64(mid)
			}else{
				cur_hour += math.Ceil(float64(dist[i]) / float64(mid))
			}
		}
		if float64(cur_hour) <= hour{
			high = mid
			res = high
		}else{
			low = mid + 1
			res = low
		}
	}
	return res
}