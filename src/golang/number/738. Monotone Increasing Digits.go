package number

//Input: N = 332
//Output: 299
func MonotoneIncreasingDigits(N int) int {
	var data []int
	for N > 0{
		rest := N % 10
		data = append([]int{rest},data...)
		N = N / 10
	}
	var l int = len(data)
	var change_pos int = l
	for i := l - 1;i > 0;i--{
		if data[i] < data[i - 1]{
			data[i - 1] = data[i - 1] - 1
			change_pos = i
		}
	}
	for i := change_pos;i < l;i++{
		data[i] = 9
	}
	var res int = 0
	for i := 0;i < l;i++{
		if data[i] == 0{
			continue
		}
		res *= 10
		res += data[i]
	}
	return res
}