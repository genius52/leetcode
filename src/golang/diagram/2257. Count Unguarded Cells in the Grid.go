package diagram

func dfs_countUnguarded(graph [][]int,rows int,columns int,r int,c int,dir []int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if graph[r][c] == 1 || graph[r][c] == 2{
		return
	}
	graph[r][c] = 3
	dfs_countUnguarded(graph,rows,columns,r + dir[0],c + dir[1],dir)
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	var graph [][]int = make([][]int,m) //0: empty, 1: guard, 2: wall,3: visited
	for i := 0;i < m;i++{
		graph[i] = make([]int,n)
	}
	for _,g := range guards{
		graph[g[0]][g[1]] = 1
	}
	for _,w := range walls{
		graph[w[0]][w[1]] = 2
	}
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,g := range guards{
		x := g[0]
		y := g[1]
		for _,dir := range dirs{
			dfs_countUnguarded(graph,m,n,x + dir[0],y + dir[1],dir)
		}
	}
	var res int = 0
	for i := 0;i < m;i++{
		for j := 0;j < n;j++{
			if graph[i][j] == 0{
				res++
			}
		}
	}
	return res
}