package diagram

func check_totalNQueens(graph [][]byte,n int,r int,c int)bool{
	var dirs [][]int = [][]int{{-1,0},{0,-1},{0,1},{-1,-1},{-1,1}}
	for _,dir := range dirs{
		var cur_r int = r
		var cur_c int = c
		for true{
			cur_r += dir[0]
			cur_c += dir[1]
			if cur_r < 0 || cur_r >= n || cur_c < 0 || cur_c >= n{
				break
			}
			if graph[cur_r][cur_c] == 'Q'{
				return false
			}
		}
	}
	return true
}

func dfs_totalNQueens(graph [][]byte,n int,r int)int{
	if r >= n{
		return 1
	}
	var res int = 0
	for i := 0;i < n;i++{
		ret := check_totalNQueens(graph,n,r,i)
		if ret{
			graph[r][i] = 'Q'
			res += dfs_totalNQueens(graph,n,r + 1)
			graph[r][i] = '.'
		}
	}
	return res
}

func totalNQueens(n int) int {
	var graph [][]byte = make([][]byte, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			graph[i][j] = '.'
		}
	}
	return dfs_totalNQueens(graph, n, 0)
}