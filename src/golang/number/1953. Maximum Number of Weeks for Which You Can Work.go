package number

func numberOfWeeks(milestones []int) int64 {
	var max_val int64 = 0
	var sum int64 = 0
	for _,m := range milestones{
		sum += int64(m)
		if int64(m) > max_val{
			max_val = int64(m)
		}
	}
	rest := sum - max_val
	if max_val > rest{
		return rest * 2 + 1
	}
	return sum
}