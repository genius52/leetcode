package number

func numberOfSteps(num int) int{
	if num == 0{
		return 0
	}
	if (num | 1) == num{
		return 1 + numberOfSteps(num ^ 1)
	}else{
		return 1 + numberOfSteps(num >> 1)
	}
}

func numberOfSteps2(num int) int{
	var steps int = 0
	for num != 0{
		if (num | 1) == num{
			num--
		}else{
			num /= 2
		}
		steps++
	}
	return steps
}