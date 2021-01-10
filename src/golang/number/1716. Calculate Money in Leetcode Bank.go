package number

func totalMoney(n int) int {
	var weeks int = n/7
	var days int = n%7
	var init_week int = 1+2+3+4+5+6+7
	var res int = init_week * weeks
	var increase int = 7
	for i := 1;i < weeks;i++{
		res += increase
		increase += 7
	}
	var start int = weeks + 1
	for i := 0;i <= days;i++{
		res += start
		start++
	}
	return res
}