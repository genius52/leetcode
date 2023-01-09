package array

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func MinimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	common := lcm(divisor1,divisor2)
	var low int = 1
	var high int = 1e9
	var res int = 1e9
	for low < high{
		var mid int = (low + high)/2
		total_cnt := mid - mid / divisor1 - mid / divisor2 + mid / common
		if total_cnt >= uniqueCnt1 + uniqueCnt2{
			res = mid
			high = mid - 1
		}else{
			low = mid + 1
		}
	}
	return res
}