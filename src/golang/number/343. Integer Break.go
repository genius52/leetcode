package number

func calc(num int,divide int)int{
	sub_num := num / divide
	rest := num % divide
	var res int = 1
	for i := 0;i < divide;i++{
		if rest > 0{
			res *= (sub_num + 1)
			rest--
		}else{
			res *= sub_num
		}
	}
	return res
}

func IntegerBreak(n int) int {
	var max_val int = 1
	for i := 2;i <= n - 1;i++{
		cur := calc(n,i)
		if cur > max_val{
			max_val = cur
		}
	}
	return max_val
}