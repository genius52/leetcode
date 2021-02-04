package number

//Input: x = -123
//Output: -321
func reverse(x int) int {
	var negative bool = false
	if x < 0{
		negative = true
		x = -x
	}
	var num int64 = 0
	for x > 0{
		num *= 10
		rest := x % 10
		num += int64(rest)
		x /= 10
	}
	if num > (1<<31 - 1) || num < -(1 << 31){
		return 0
	}
	if negative{
		return int(-num)
	}
	return int(num)
}