package number

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var res int = 0
	for _, h := range hours {
		if h >= target {
			res++
		}
	}
	return res
}
