package number

func MinimizeXor(num1 int, num2 int) int {
	var record [32]bool
	var n1 int = num1
	var offset int = 0
	for n1 > 0{
		if n1 | 1 == n1{
			record[offset] = true
		}
		offset++
		n1 >>= 1
	}
	var one_cnt int = 0
	var n2 int = num2
	for n2 > 0{
		if n2 | 1 == n2{
			one_cnt++
		}
		n2 >>= 1
	}
	var res int = 0
	for i := 31;i >= 0 && one_cnt > 0;i--{
		if record[i]{
			res |= 1 << i
			one_cnt--
		}
	}
	for i := 0;i < 32 && one_cnt > 0;i++{
		if !record[i]{
			res |= 1 << i
			one_cnt--
		}
	}
	return res
}