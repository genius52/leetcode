package array

//Input: grid = [[0,1,-1],[1,0,-1],[1,1,1]]
//Output: 5

//Input: grid = [[1,1,-1],[1,-1,1],[-1,1,1]]
//Output: 0
func dp_cherryPickup(grid [][]int,rows int,columns int,r1 int,c1 int,r2 int,c2 int,memo [][][][]int)int{
	if r1 == (rows - 1) && c1 == (columns - 1){
		return grid[r1][c1]
	}
	if r2 == (rows - 1) && c2 == (columns - 1){
		return grid[r2][c2]
	}

	if r1 < 0 || r1 >= rows || c1 < 0 || c1 >= columns || r2 < 0 || r2 >= rows || c2 < 0 || c2 >= columns{
		return -2147483648
	}
	if grid[r1][c1] == -1 || grid[r2][c2] == -1{
		return -2147483648
	}
	if memo[r1][c1][r2][c2] != 0 {
		return memo[r1][c1][r2][c2]
	}
	var res int = 0
	if r1 == r2 && c1 == c2{
		res = grid[r1][c1]
	}else{
		res = grid[r1][c1] + grid[r2][c2]
	}
	memo[r1][c1][r2][c2] = res + max_int_number(dp_cherryPickup(grid,rows,columns,r1 + 1,c1,r2 + 1,c2,memo),
		dp_cherryPickup(grid,rows,columns,r1 + 1,c1,r2,c2 + 1,memo),
		dp_cherryPickup(grid,rows,columns,r1,c1 + 1,r2 + 1,c2,memo),
		dp_cherryPickup(grid,rows,columns,r1,c1 + 1,r2,c2 + 1,memo))
	return memo[r1][c1][r2][c2]
}

func CherryPickup(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var memo [][][][]int = make([][][][]int,rows)
	for i := 0;i < rows;i++{
		memo[i] = make([][][]int,columns)
		for j := 0;j < columns;j++{
			memo[i][j] = make([][]int,rows)
			for k := 0;k < rows;k++{
				memo[i][j][k] = make([]int,columns)
			}
		}
	}
	res := dp_cherryPickup(grid,rows,columns,0,0,0,0,memo)
	if res < 0{
		return 0
	}
	return res
}