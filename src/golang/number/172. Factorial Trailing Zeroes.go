package number

func count_five(n int)int{
	var res int = 0
	for n > 0{
		if n % 5 != 0{
			return res
		}
		res++
		n /= 5
	}
	return res
}

func trailingZeroes(n int) int {
	var res int = 0
	for n > 0{
		res += count_five(n)
		if res > 0{
			n -= 5
		}else{
			n--
		}
	}
	return res
}