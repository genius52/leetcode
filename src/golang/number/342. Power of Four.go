package number

func IsPowerOfFour(n int) bool {
	if n <= 0{
		return false
	}
	var mask int64 = 1
	var zero_cnt int = 0
	for mask < 2147483647{
		if (int64(n) | mask) == int64(n){
			break
		}
		zero_cnt++
		mask <<= 1
	}
	if (zero_cnt | 1) == zero_cnt{
		return false
	}
	mask <<= 1
	for mask < 2147483647{
		if (int64(n) | mask) == int64(n){
			return false
		}
		mask <<= 1
	}
	return true
}