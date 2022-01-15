package number

func count_countGoodNumbers(n int,l int64)int64{
	var res int64 = 1
	var mul int64 = int64(n)
	for l > 0{
		if l % 2 == 1{
			res *= mul
			res %= 1e9 + 7
		}
		mul = mul * mul
		mul %= 1e9 + 7
		l /= 2
	}
	return res
}

func CountGoodNumbers(n int64) int {
	return int(count_countGoodNumbers(5,(n + 1)/2) * count_countGoodNumbers(4,n/2) % (1e9 + 7))
}
