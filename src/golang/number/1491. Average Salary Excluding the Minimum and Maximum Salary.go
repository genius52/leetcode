package number

func average(salary []int) float64 {
	var l int = len(salary)
	if l <= 2{
		return 0
	}
	var sum int = 0
	var min_val int = 2147483647
	var max_val int = 0
	for i := 0;i < l;i++{
		sum += salary[i]
		if salary[i] > max_val{
			max_val = salary[i]
		}
		if salary[i] < min_val{
			min_val = salary[i]
		}
	}
	return float64(sum - min_val - max_val)/float64(l - 2)
}