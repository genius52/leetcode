package array

func matrixScore(A [][]int) int {
	//reverse line
	for i := 0;i < len(A); i++{
		if A[i][0] != 1{
			for j := 0;j < len(A[i]) ; j++{
				if A[i][j] == 0{
					A[i][j] = 1
				}else{
					A[i][j] = 0
				}
			}
		}
	}
	//reverse column
	for i := 0; i < len(A[0]); i++{
		var zero_cnt int = 0
		for j := 0 ; j < len(A);j++{
			if A[j][i] == 0{
				zero_cnt++
			}
		}
		if zero_cnt > len(A)/2{
			for j := 0 ; j < len(A);j++{
				if A[j][i] == 0{
					A[j][i] = 1
				}else{
					A[j][i] = 0
				}
			}
		}
	}
	//calculate
	var res int = 0
	for i := 0;i < len(A); i++{
		var val int = 0
		for j := len(A[i]) - 1;j >= 0 ; j-- {
			val += A[i][j] << uint32(len(A[i]) - 1 - j)
		}
		res += val
	}
	return res
}