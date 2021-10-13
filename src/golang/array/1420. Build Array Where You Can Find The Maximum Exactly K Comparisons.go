package array

//arr 中有 n 个整数。
//1 <= arr[i] <= m 其中 (0 <= i < n) 。
//将上面提到的算法应用于 arr ，search_cost 的值等于 k

//输入：n = 2, m = 3, k = 1
//输出：6
//解释：可能的数组分别为 [1, 1], [2, 1], [2, 2], [3, 1], [3, 2] [3, 3]
func dp_numOfArrays(n int,m int,k int,cur_cnt int,cur_max int,cur_cost int,memo *[][][]int)int{
	if cur_max > m{
		return 0
	}
	if cur_cnt == n{

	}
	max_val := -1
	//max_idx := -1
	cost := 0
	var res int = 0
	for i := 0;i < cur_max;i++{
		if max_val < i{
			max_val = i
			cost++
		}
		if cost > m{
			break
		}
	}
	return res
}

func numOfArrays(n int, m int, k int) int {
	var memo [][][]int = make([][][]int,n + 1)
	for i := 0;i <= n;i++{
		memo[i] = make([][]int,m + 1)
		for j := 0;j <= m;j++{
			memo[i][j] = make([]int,k + 1)
		}
	}
	return dp_numOfArrays(n,m,k,0,0,0,&memo)
}