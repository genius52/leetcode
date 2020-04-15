package number

import "math"

func AngleClock(hour int, minutes int) float64 {
	var hour_degree float64 = float64(hour) * 30 + float64(minutes) * 0.5
	var minutes_degree float64 = float64(minutes) * 6
	res := math.Abs(hour_degree  - minutes_degree)
	if res > 180{
		return (360 - res)
	}else{
		return res
	}
}