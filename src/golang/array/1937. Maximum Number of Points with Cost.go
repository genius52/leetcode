package array

func MaxPoints(points [][]int) int64 {
	var rows int = len(points)
	var columns int = len(points[0])
	var pre []int64 = make([]int64,columns)
	for j := 0;j < columns;j++{
		pre[j] = int64(points[0][j])
	}
	for i := 1;i < rows;i++{
		var left_max int64 = pre[0]
		var cur_left []int64 = make([]int64,columns)
		cur_left[0] = int64(points[i][0]) + pre[0]
		for j := 1;j < columns;j++{
			cur_left[j] = int64(points[i][j]) + max_int64(int64(pre[j]),left_max - 1)
			left_max = max_int64(left_max - 1,int64(pre[j]))
		}
		var right_max int64 = pre[columns - 1]
		var cur_right []int64 = make([]int64,columns)
		cur_right[columns - 1] = int64(points[i][columns - 1]) + pre[columns - 1]
		for j := columns - 2;j >= 0;j--{
			cur_right[j] = int64(points[i][j]) + max_int64(int64(pre[j]),right_max - 1)
			right_max = max_int64(right_max - 1,int64(pre[j]))
		}
		var cur []int64 = make([]int64,columns)
		for j := 0;j < columns;j++{
			cur[j] = max_int64(cur_left[j],cur_right[j])
		}
		//copy(pre,cur)
		pre = cur
	}
	var res int64 = 0
	for j := 0;j < columns;j++{
		res = max_int64(res,pre[j])
	}
	return res
}