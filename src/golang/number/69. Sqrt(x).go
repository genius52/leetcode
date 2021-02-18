package number


func mySqrt(x int) int {
	if x == 0{
		return 0
	}
	low := 1
	high := x
	for low <= high{
		mid := low + (high-low)/2
		val := mid*mid
		if val == x{
			return mid
		}else if val < x{
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return high
}
