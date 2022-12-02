package number

func pivotInteger(n int) int {
	var left int = 1
	var right int = n
	for left <= right{
		var mid int = (left + right)/2
		sum1 := (1 + mid) * mid/2
		sum2 := (mid + n) * (n - mid + 1)/2
		if sum1 == sum2{
			return mid
		}else if sum1 < sum2{
			left++
		}else if sum1 > sum2{
			right--
		}
	}
	return -1
}