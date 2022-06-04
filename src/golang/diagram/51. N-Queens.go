package diagram

import "strings"

func check_solveNQueens(graph [][]byte,n int,r int,c int)bool{
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

func dfs_solveNQueens(graph [][]byte,n int,r int,res *[][]string){
	if r >= n{
		var s []string
		for i := 0;i < n;i++{
			var ss strings.Builder
			for j := 0;j < n;j++{
				ss.WriteByte(graph[i][j])
			}
			s = append(s,ss.String())
		}
		*res = append(*res,s)
		return
	}
	for i := 0;i < n;i++{
		ret := check_solveNQueens(graph,n,r,i)
		if ret{
			graph[r][i] = 'Q'
			dfs_solveNQueens(graph,n,r + 1,res)
			graph[r][i] = '.'
		}
	}
}

func SolveNQueens(n int) [][]string {
	var res [][]string
	var graph [][]byte = make([][]byte, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			graph[i][j] = '.'
		}
	}
	dfs_solveNQueens(graph, n, 0, &res)
	return res
}