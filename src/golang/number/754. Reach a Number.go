package number

func ReachNumber(target int) int {
	if target < 0{
		target = -target
	}
	var each_step int = 1
	var total int = 0
	var steps int = 0
	for total < target{
		total += each_step
		each_step++
		steps++
	}
	if (total - target) % 2 == 0{
		return steps
	}
	for (total - target) % 2 != 0{
		steps++
		total += each_step
		each_step++
	}
	return steps
}