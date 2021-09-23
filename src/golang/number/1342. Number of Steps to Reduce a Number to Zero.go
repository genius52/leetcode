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