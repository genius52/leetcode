package diagram

//Input: m = 2, n = 2, N = 2, i = 0, j = 0
//Output: 6
func dfs_findPaths(m int, n int, move int, startRow int, startColumn int,memo [][][]int)int{
	if startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n{
		return 1
	}
	if move == 0{
		return 0
	}
	if memo[startRow][startColumn][move] != -1{
		return memo[startRow][startColumn][move]
	}
	var res int = 0
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var MOD int = 1e9 + 7
	for _,dir := range dirs{
		res += dfs_findPaths(m,n,move - 1,startRow + dir[0],startColumn + dir[1],memo)
		res %= MOD
	}
	memo[startRow][startColumn][move] = res
	return res
}

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var memo [][][]int = make([][][]int,m)
	for i := 0;i < m;i++{
		memo[i] = make([][]int,n)
		for j := 0;j < n;j++{
			memo[i][j] = make([]int,maxMove + 1)
			for k := 0;k <= maxMove;k++{
				memo[i][j][k] = -1
			}
		}
	}
	return dfs_findPaths(m,n,maxMove,startRow,startColumn,memo)
}

type Point struct{
	x int
	y int
}

type Point2 struct{
	x int
	y int
	left_step int
}

//func dfs_findPaths(i int,j int,rows int,columns int,cur_steps int,N int,record map[Point2]int) int{
//	if i < 0 || i >= rows || j < 0 || j >= columns{
//		return 1
//	}
//	if cur_steps == N{
//		return 0
//	}
//	rest_step := N - cur_steps
//	if ((i - rest_step) >= 0 && (i + rest_step) < rows) && ((j - rest_step) >= 0 && (j + rest_step) < columns){
//		return 0
//	}
//	var p Point2 = Point2{i,j,rest_step}
//	if _,ok := record[p];ok{
//		return record[p]
//	}
//	res := (dfs_findPaths(i - 1,j,rows,columns,cur_steps + 1,N,record) + dfs_findPaths(i + 1,j,rows,columns,cur_steps + 1,N,record) +
//		dfs_findPaths(i,j - 1,rows,columns,cur_steps + 1,N,record) + dfs_findPaths(i,j + 1,rows,columns,cur_steps + 1,N,record)) % (1e9 + 7)
//	record[p] = res
//	return res
//}
//
//func FindPaths(m int, n int, N int, i int, j int) int {
//	var record map[Point2]int = make(map[Point2]int)
//	return dfs_findPaths(i,j,m,n,0,N,record)
//}