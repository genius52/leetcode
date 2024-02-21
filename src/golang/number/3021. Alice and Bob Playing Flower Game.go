package number

func FowerGame(n int, m int) int64 {
	//Alice win: n + m == odd
	var n_even int64 = int64(n/2) * int64(m/2+m%2)
	var n_odd int64 = int64(n/2+n%2) * int64(m/2)
	return n_even + n_odd
}
