package number

func isPerfectSquare(num int) bool {
	var left int = 1
	var right int = num
	for left <= right{
		mid := (left + right)/2
		val := mid * mid
		if val == num{
			return true
		}else if val < num{
			left = mid + 1
		}else if val > num{
			right = mid - 1
		}
	}
	return false
}

func isPerfectSquare2(num int) bool{
	var cur int = 1
	var sum int = 0
	for sum < num{
		sum += cur
		if sum == num{
			return true
		}
		cur += 2
	}
	return false
}