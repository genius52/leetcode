package array

import "math"

func MaxArea(height []int) int {
	var length int = len(height)
	var left int = 0
	var right int = length - 1
	var max_cap int = 0
	for left < right  {
		var cap int = int(math.Abs(float64(right - left)) * math.Min(float64(height[right]),float64(height[left])))
		if cap > max_cap{
			max_cap = cap
		}
		if height[right] < height[left]{
			right--
		} else{
			left++
		}
	}
	return max_cap
}