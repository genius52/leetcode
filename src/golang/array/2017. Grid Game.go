package array

//Input: grid = [[1,3,1,15],[1,3,3,1]]
//Output: 7
//func max_int64(a,b int64)int64{
//	if a > b {
//		return a
//	}else{
//		return b
//	}
//}
func gridGame(grid [][]int) int64 {
	var columns int = len(grid[0])
	var prefix_sum []int64 = make([]int64,columns)
	var suffix_sum []int64 = make([]int64,columns)
	prefix_sum[0] = int64(grid[0][0])
	for i := 1;i < columns;i++{
		prefix_sum[i] = prefix_sum[i - 1] + int64(grid[0][i])
	}
	suffix_sum[columns - 1] = int64(grid[1][columns - 1])
	for i := columns - 2;i >= 0;i--{
		suffix_sum[i] = suffix_sum[i + 1] + int64(grid[1][i])
	}
	var max_score2 int64 = 1 << 62
	for i := 0;i < columns;i++{
		cur := max_int64(prefix_sum[columns - 1] - prefix_sum[i],suffix_sum[0] - suffix_sum[i])
		if cur < max_score2{
			max_score2 = cur
		}
	}
	return max_score2
}