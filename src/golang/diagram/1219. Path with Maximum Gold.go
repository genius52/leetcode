package diagram

func dfs_getMaximumGold(grid [][]int,rows int,columns int,r int,c int,visited [][]bool)int{
	if r < 0 || r >= rows || c < 0 || c >= columns || grid[r][c] == 0 || visited[r][c]{
		return 0
	}
	visited[r][c] = true
	var res int = grid[r][c] + max_int_number(dfs_getMaximumGold(grid,rows,columns,r - 1,c,visited),
		dfs_getMaximumGold(grid,rows,columns,r + 1,c,visited),
		dfs_getMaximumGold(grid,rows,columns,r,c - 1,visited),
		dfs_getMaximumGold(grid,rows,columns,r,c + 1,visited))
	visited[r][c] = false
	return res
}

func GetMaximumGold(grid [][]int) int{
	var rows int = len(grid)
	var columns int = len(grid[0])

	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			var visited [][]bool = make([][]bool,rows)
			for i := 0;i < rows;i++{
				visited[i] = make([]bool,columns)
			}
			if grid[i][j] == 0{
				continue
			}
			cur := dfs_getMaximumGold(grid,rows,columns,i,j,visited)
			if cur > res{
				res = cur
			}
		}
	}
	return res
}

func dfs_gold(grid [][]int,x int,y int,record [][]int)int{
	if x >= len(grid) || y >= len(grid[0]) || x < 0 || y < 0{
		return 0
	}
	if grid[x][y] == 0 || record[x][y] != 0{
		return 0
	}
	record[x][y] = 1
	var dup_record1,dup_record2,dup_record3,dup_record4 [][]int = make([][]int,len(grid)),make([][]int,len(grid)),make([][]int,len(grid)),make([][]int,len(grid))
	for i := 0;i < len(grid);i++ {
		dup_record1[i] = make([]int,len(grid[i]))
		dup_record2[i] = make([]int,len(grid[i]))
		dup_record3[i] = make([]int,len(grid[i]))
		dup_record4[i] = make([]int,len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			dup_record1[i][j] = record[i][j]
			dup_record2[i][j] = record[i][j]
			dup_record3[i][j] = record[i][j]
			dup_record4[i][j] = record[i][j]
		}
	}
	var res int = 0
	res = max_int(grid[x][y] + dfs_gold(grid,x - 1,y,dup_record1),grid[x][y] + dfs_gold(grid,x,y-1,dup_record2))
	res = max_int(res,grid[x][y] + dfs_gold(grid,x + 1,y,dup_record3))
	res = max_int(res,grid[x][y] + dfs_gold(grid,x,y + 1,dup_record4))
	return res
}

func getMaximumGold(grid [][]int) int {
	var max int = 0
	for i := 0;i < len(grid);i++{
		for j := 0;j < len(grid[i]);j++{
			if grid[i][j] != 0{
				var record [][]int = make([][]int,len(grid))
				for row := 0; row < len(grid);row++{
					record[row] = make([]int,len(grid[i]))
				}
				max = max_int(max,dfs_gold(grid,i,j,record))
			}
		}
	}
	return max
}