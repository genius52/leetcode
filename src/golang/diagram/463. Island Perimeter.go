package diagram

//Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
//Output: 16
//Explanation: The perimeter is the 16 yellow stripes in the image above.
func dfs_islandPerimeter(grid [][]int,rows int,columns int,r int,c int)int{
	if r < 0 || c < 0 || r >= rows || c >= columns || grid[r][c] == 0{
		return -1
	}
	if grid[r][c] == 2{
		return 0
	}
	grid[r][c] = 2
	var total int = 0
	res := dfs_islandPerimeter(grid,rows,columns,r - 1,c)
	if res == -1{
		total++
	}else if res > 0{
		total += res
	}
	res = dfs_islandPerimeter(grid,rows,columns,r + 1,c)
	if res == -1{
		total++
	}else if res > 0{
		total += res
	}
	res = dfs_islandPerimeter(grid,rows,columns,r,c - 1)
	if res == -1{
		total++
	}else if res > 0{
		total += res
	}
	res = dfs_islandPerimeter(grid,rows,columns,r,c + 1)
	if res == -1{
		total++
	}else if res > 0{
		total += res
	}
	return total
}

func islandPerimeter(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 0{
				continue
			}
			res = dfs_islandPerimeter(grid,rows,columns,i,j)
			return res
		}
	}
	return res
}