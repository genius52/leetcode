package diagram

func dfs_hitBricks(grid [][]int,rows int,columns int,r int,c int,nodrop map[int]bool){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if grid[r][c] == 0{
		return
	}
	index := r * columns + c
	if _,ok := nodrop[index];ok{
		return
	}
	nodrop[index] = true
	dfs_hitBricks(grid,rows,columns,r - 1,c,nodrop)
	dfs_hitBricks(grid,rows,columns,r + 1,c,nodrop)
	dfs_hitBricks(grid,rows,columns,r,c - 1,nodrop)
	dfs_hitBricks(grid,rows,columns,r,c + 1,nodrop)
}

func HitBricks(grid [][]int, hits [][]int) []int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	//var nodrop map[int]bool = make(map[int]bool)
	var total int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 1{
				total++
				//nodrop[i * columns + j] = true
			}
		}
	}
	var l int = len(hits)
	var res []int = make([]int, l)
	for i := 0;i < l;i++{
		if grid[hits[i][0]][hits[i][1]] == 0{
			continue
		}
		grid[hits[i][0]][hits[i][1]] = 0
		total--
		var cur_nodrop map[int]bool = make(map[int]bool)
		for j := 0;j < columns;j++{
			dfs_hitBricks(grid,rows,columns,0,j,cur_nodrop)
		}
		res[i] = total - len(cur_nodrop)
		total = len(cur_nodrop)
	}
	return res
}