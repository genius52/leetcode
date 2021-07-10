package array

func MinFallingPathSum(A [][]int) int{
	var rows int = len(A)
	var columns int = len(A[0])
	var pre []int = make([]int,columns)

	copy(pre,A[0])
	for i := 1;i < rows;i++{
		var cur []int = make([]int,columns)
		for j := 0;j < columns;j++{
			if j == 0{
				cur[j] = A[i][j] + min_int_number(pre[j],pre[j + 1])
			}else if j == (columns - 1){
				cur[j] = A[i][j] + min_int_number(pre[j - 1],pre[j])
			}else{
				cur[j] = A[i][j] + min_int_number(pre[j - 1],pre[j],pre[j + 1])
			}
		}
		pre = cur
	}
	var res int = 2147483647
	for i := 0;i < columns;i++{
		if pre[i] < res{
			res = pre[i]
		}
	}
	return res
}

//func dp_minpath(A [][]int,i int,j int,memo [][]int)int{
//	if i >= len(A) || i < 0 || j >= len(A[0]) || j < 0 {
//		return 0
//	}
//	if memo[i][j] != 0{
//		return memo[i][j]
//	}
//	var res int = math.MaxInt32
//	res = int(math.Min(float64(A[i][j] + dp_minpath(A,i+1,j,memo)),float64(res)))
//	if (j + 1) < len(A[0]) {
//		res = int(math.Min(float64(A[i][j]+dp_minpath(A, i+1, j+1,memo)), float64(res)))
//	}
//	if (j - 1) >= 0 {
//		res = int(math.Min(float64(A[i][j]+dp_minpath(A, i+1, j-1,memo)), float64(res)))
//	}
//	memo[i][j] = res
//	return res
//}
//
//func minFallingPathSum(A [][]int) int {
//	var res int
//	var memo [][]int = make([][]int,len(A))
//	for i := 0; i < len(A);i++{
//		memo[i] = make([]int,len(A[0]))
//	}
//	for j := 0; j < len(A[0]);j++{
//		res = int(math.Min(float64(dp_minpath(A,0,j,memo)),float64(math.MaxInt32)))
//	}
//	for _,v := range memo[0]{
//		if v < res{
//			res = v
//		}
//	}
//	return res
//}