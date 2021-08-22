package number


func tribonacci(n int) int {
	if n <= 0{
		return 0
	}
	if n == 1 || n == 2{
		return 1
	}
	n1 := 0
	n2 := 1
	n3 := 1
	res := 0
	for i := 3;i <= n;i++{
		res = n3 + n2 + n1
		n1 = n2
		n2 = n3
		n3 = res
	}
	return res
}


//1137
func dp_tribonacci(n int,record map[int]int)int {
	if n <= 0{
		return 0
	}
	if n == 1 || n == 2{
		return 1
	}

	if _,ok := record[n];ok{
		return record[n]
	}
	sum := 0
	for i := 1;i <= 3 ;i++{
		sum += tribonacci(n - i)
	}
	record[n] = sum
	return sum
}
