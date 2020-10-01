package number

//Input: N = 332
//Output: 299
func MonotoneIncreasingDigits(N int) int {
	var data []int
	for N > 0{
		rest := N % 10
		data = append(data,rest)
		N = N / 10
	}
	var l int = len(data)
	for i := 0;i < l/2;i++{
		data[i],data[l - 1 - i] = data[l - 1 - i],data[i]
	}
	var pos int = l
	for i := l - 1;i > 0;i--{
		if data[i] < data[i - 1]{
			data[i - 1] = data[i - 1] - 1
			pos = i
		}
	}
	for i := pos;i < l;i++{
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