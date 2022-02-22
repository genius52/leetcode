package number

func countEven(num int) int {
	var sum int = 0
	var n int = num
	for n > 0{
		sum += n % 10
		n /= 10
	}
	var res int = 0
	if (sum | 1) != sum{
		res = num/2
	}else{
		res = (num - 1) /2
	}
	return res
}