package number

func rand7()int{
	return 5
}

func rand10() int {
	var res int = 0
	for i := 0;i < 10;i++{
		res += rand7()
	}
	return res % 10 + 1
}