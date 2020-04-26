package diagram

//Input: m = 2, n = 2, N = 2, i = 0, j = 0
//Output: 6
type Point struct{
	x int
	y int
	left_step int
}
func dfs_findPaths(i int,j int,rows int,columns int,cur_steps int,N int,record map[Point]int) int{
	if i < 0 || i >= rows || j < 0 || j >= columns{
		return 1
	}
	if cur_steps == N{
		return 0
	}
	rest_step := N - cur_steps
	if ((i - rest_step) >= 0 && (i + rest_step) < rows) && ((j - rest_step) >= 0 && (j + rest_step) < columns){
		return 0
	}
	var p Point = Point{i,j,rest_step}
	if _,ok := record[p];ok{
		return record[p]
	}
	res := (dfs_findPaths(i - 1,j,rows,columns,cur_steps + 1,N,record) + dfs_findPaths(i + 1,j,rows,columns,cur_steps + 1,N,record) +
		dfs_findPaths(i,j - 1,rows,columns,cur_steps + 1,N,record) + dfs_findPaths(i,j + 1,rows,columns,cur_steps + 1,N,record)) % (1e9 + 7)
	record[p] = res
	return res
}

func FindPaths(m int, n int, N int, i int, j int) int {
	var record map[Point]int = make(map[Point]int)
	return dfs_findPaths(i,j,m,n,0,N,record)
}