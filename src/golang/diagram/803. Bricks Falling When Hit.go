package diagram

func dfs_calc_stable(grid [][]int,rows int,columns int,r int,c int,tag int)int{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return 0
	}
	if grid[r][c] == -1 || grid[r][c] == 0 || grid[r][c] == tag{
		return 0
	}
	grid[r][c] = tag
	return 1 + dfs_calc_stable(grid,rows,columns,r - 1,c,tag) +
		dfs_calc_stable(grid,rows,columns,r + 1,c,tag) +
		dfs_calc_stable(grid,rows,columns,r,c - 1,tag) +
		dfs_calc_stable(grid,rows,columns,r,c + 1,tag)
}

func is_stable_around(grid [][]int, rows int,columns int,r int,c int)bool  {
	if r == 0{
		return true
	}
	dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,dir := range dirs{
		next_r := r + dir[0]
		next_c := c + dir[1]
		if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns{
			continue
		}
		if grid[next_r][next_c] == 2{
			return true
		}
	}
	return false
}

func HitBricks(grid [][]int, hits [][]int) []int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	var l int = len(hits)
	for i := 0;i < l;i++{
		grid[hits[i][0]][hits[i][1]]-- //标记被击落的位置，如果原本没有砖块，则没有任何影响
	}
	for j := 0;j < columns;j++{
		dfs_calc_stable(grid,rows,columns,0,j,2)//最终不会掉落的设置成2
	}
	var res []int = make([]int,l)
	for i := l - 1;i >= 0;i--{
		r := hits[i][0]
		c := hits[i][1]
		grid[r][c]++
		if !is_stable_around(grid,rows,columns,r,c) || grid[r][c] == 0{
			res[i] = 0
		}else{//周围有稳定砖块时，才存在击落砖块的情况，否则不会击落任何砖块
			//计算能连接到值为1的砖块
			res[i] = dfs_calc_stable(grid,rows,columns,r,c,2) - 1
		}
	}
	return res
}

//func dfs_hitBricks(grid [][]int,rows int,columns int,r int,c int,nodrop map[int]bool){
//	if r < 0 || r >= rows || c < 0 || c >= columns{
//		return
//	}
//	if grid[r][c] == 0{
//		return
//	}
//	index := r * columns + c
//	if _,ok := nodrop[index];ok{
//		return
//	}
//	nodrop[index] = true
//	dfs_hitBricks(grid,rows,columns,r - 1,c,nodrop)
//	dfs_hitBricks(grid,rows,columns,r + 1,c,nodrop)
//	dfs_hitBricks(grid,rows,columns,r,c - 1,nodrop)
//	dfs_hitBricks(grid,rows,columns,r,c + 1,nodrop)
//}
//
//func HitBricks(grid [][]int, hits [][]int) []int {
//	var rows int = len(grid)
//	var columns int = len(grid[0])
//	//var nodrop map[int]bool = make(map[int]bool)
//	var total int = 0
//	for i := 0;i < rows;i++{
//		for j := 0;j < columns;j++{
//			if grid[i][j] == 1{
//				total++
//				//nodrop[i * columns + j] = true
//			}
//		}
//	}
//	var l int = len(hits)
//	var res []int = make([]int, l)
//	for i := 0;i < l;i++{
//		if grid[hits[i][0]][hits[i][1]] == 0{
//			continue
//		}
//		grid[hits[i][0]][hits[i][1]] = 0
//		total--
//		var cur_nodrop map[int]bool = make(map[int]bool)
//		for j := 0;j < columns;j++{
//			dfs_hitBricks(grid,rows,columns,0,j,cur_nodrop)
//		}
//		res[i] = total - len(cur_nodrop)
//		total = len(cur_nodrop)
//	}
//	return res
//}