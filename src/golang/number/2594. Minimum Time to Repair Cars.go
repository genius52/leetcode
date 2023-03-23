package number

import (
	"math"
)

// 假设最少的时间是minutes，能否完成cars的维修？
func check_repairCars(ranks []int, cars int, minutes int64) bool {
	var res int = 0
	for _, val := range ranks {
		res += int(math.Floor(math.Sqrt(float64(minutes / int64(val)))))
	}
	return res >= cars
}

func repairCars(ranks []int, cars int) int64 {
	//var l int = len(ranks)
	var low int64 = 1
	var high int64 = 0
	for _, val := range ranks {
		var cur int64 = int64(val * cars * cars)
		if cur > high {
			high = cur
		}
	}
	for low < high {
		mid := (low + high) / 2
		ret := check_repairCars(ranks, cars, mid)
		if ret {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
