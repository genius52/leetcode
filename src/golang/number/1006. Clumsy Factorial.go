package number

func recursive_clumsy(N int,pre int,res *int){
	if N == 1{
		if pre == 0{
			*res = 1
		}else{
			*res = pre - 1
		}
	}else if N == 2{
		if pre == 0{
			*res = 2
		}else{
			*res = pre - 2 * 1
		}
	}else if N == 3{
		if pre == 0{
			*res = 3 * 2 / 1
		}else {
			*res = pre - 3 * 2 / 1
		}
	}else if N == 4{
		if pre == 0{
			*res = 4 * 3 / 2 + 1
		}else {
			*res = pre - 4 * 3 / 2 + 1
		}
	}else{
		if pre == 0{
			pre = N * (N - 1) / (N - 2) + (N - 3)
		}else{
			pre = pre - N * (N - 1) / (N - 2) + (N - 3)
		}
		recursive_clumsy(N - 4,pre,res)
	}
}

func Clumsy(N int) int {
	var res int = 0
	recursive_clumsy(N,0,&res)
	return res
}