package number

func IsPowerOfTwo(n int) bool {
	if n <= 0{
		return false
	}
	var mask int = 1
	var find_one bool = false
	for i := 0;i < 32;i++{
		if (n & mask) == mask{
			if find_one{
				return false
			}
			find_one = true
		}
		mask <<= 1
	}
	return true
}

func isPowerOfTwo(n int) bool{
	if n <= 0{
		return false
	}
	return (n & (n - 1)) == 0
}