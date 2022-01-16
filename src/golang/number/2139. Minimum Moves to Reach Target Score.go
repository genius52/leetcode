package number

func recursive_minMoves(n int,maxDoubles int)int{
	if n == 1{
		return 0
	}
	if maxDoubles == 0{
		return n - 1
	}
	if n | 1 == n{//奇数
		return 1 + recursive_minMoves(n - 1,maxDoubles)
	}else{
		return 1 + recursive_minMoves(n/2,maxDoubles - 1)
	}
}

func minMoves2139(target int, maxDoubles int) int {
	if maxDoubles == 0{
		return target - 1
	}
	return recursive_minMoves(target,maxDoubles)
}