package diagram

func maximalNetworkRank(n int, roads [][]int) int {
	var cnt []int = make([]int,n)
	var graph [][]bool = make([][]bool,n)
	for i := 0;i < n;i++{
		graph[i] = make([]bool,n)
	}
	for _,r := range roads{
		cnt[r[0]]++
		cnt[r[1]]++
		graph[r[0]][r[1]] = true
		graph[r[1]][r[0]] = true
	}
	var res int = 0
	for i := 0;i < n;i++{
		for j := i + 1;j < n;j++{
			cur_cnt := cnt[i] + cnt[j]
			if graph[i][j]{
				cur_cnt--
			}
			if cur_cnt > res{
				res = cur_cnt
			}
		}
	}
	return res
}