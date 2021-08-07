package number

func dp_minScoreTriangulation(A []int,begin int,end int,memo *[50][50]int)int{
	if memo[begin][end] != 0{
		return memo[begin][end]
	}
	if (end - begin) < 2{
		return 0
	}
	var res int = 2147483647
	for i := begin + 1;i < end;i++{
		cur := A[begin] * A[i] * A[end] + dp_minScoreTriangulation(A,begin,i,memo) +
			dp_minScoreTriangulation(A,i,end,memo)
		if cur < res{
			res = cur
		}
	}
	memo[begin][end] = res
	return res
}

func MinScoreTriangulation(A []int) int{
	var memo [50][50]int
	return dp_minScoreTriangulation(A,0,len(A) - 1,&memo)
}

//func MinScoreTriangulation(A []int) int {
//	var l int = len(A)
//	sum := 0
//	for i := 0;i < l;i++{
//		sum += (A[i] * A[(i + 1) % l])
//	}
//	var res int = 2147483647
//	for i := 0;i < l;i++{
//		dup := A[i] * A[i] * A[(i + 1) % l] + A[i] * A[i] * A[(i + l - 1) % l]
//		cur := A[i] * sum - dup
//		if cur < res{
//			res = cur
//		}
//	}
//	return res
//}