package number

//row 1: 0
//row 2: 01
//row 3: 0110
//row 4: 01101001
//cur N + k = last k/2 + k % 2
func recursive_kthGrammar(N int,K int)int{
	if N == 1{
		return 0
	}
	last := recursive_kthGrammar(N - 1,(K + 1)/2)
	if last == 0{
		if K % 2 == 0{
			return  1
		}else{
			return 0
		}
	}else{
		if K % 2 == 0{
			return 0
		}else{
			return 1
		}
	}
}

func KthGrammar(N int, K int) int {
	return recursive_kthGrammar(N,K)
}