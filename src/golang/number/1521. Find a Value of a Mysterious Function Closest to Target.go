package number

//输入：arr = [9,12,3,7,15], target = 5
//输出：2
//解释：所有可能的 [l,r] 数对包括 [[0,0],[1,1],[2,2],[3,3],[4,4],[0,1],[1,2],[2,3],[3,4],[0,2],[1,3],[2,4],[0,3],[1,4],[0,4]]，
//Winston 得到的相应结果为 [9,12,3,7,15,8,0,3,7,0,0,3,0,0,0] 。最接近 5 的值是 7 和 3，所以最小差值为 2 。
func abs_int(n int)int{
	if n < 0{
		return -n
	}
	return n
}

func ClosestToTarget(arr []int, target int) int {
	var l int = len(arr)
	var record []map[int]bool = make([]map[int]bool,l)
	for i := 0;i < l;i++{
		record[i] = make(map[int]bool)
	}
	record[l - 1][arr[l - 1]] = true
	var res int = abs_int(arr[l - 1] - target)
	for i := l - 2;i >= 0;i--{
		record[i][arr[i]] = true
		res = min_int(res,abs_int(arr[i] - target))
		for val,_ := range record[i + 1]{
			res = min_int(res,abs_int(arr[i] & val - target))
			record[i][arr[i] & val] = true
		}
	}
	return res
}