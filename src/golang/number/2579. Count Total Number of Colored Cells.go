package number

func coloredCells(n int) int64 {
	if n == 1 {
		return 1
	}
	var res int64 = int64(n*2 - 1)
	for i := 1; i < n*2-1; i += 2 {
		res += int64(i * 2)
	}
	return res
}
