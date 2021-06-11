package diagram

func calc_largestIsland(grid [][]int,rows int,columns int,r int,c int)int{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return 0
	}
	if grid[r][c] == 0 || grid[r][c] == -1{
		return 0
	}
	grid[r][c] = -1
	return 1 + calc_largestIsland(grid,rows,columns,r - 1,c) + calc_largestIsland(grid,rows,columns,r + 1,c) +
		calc_largestIsland(grid,rows,columns,r,c - 1) + calc_largestIsland(grid,rows,columns,r,c + 1)
}

func mark_largestIsland(grid [][]int,groups [][]Group,rows int,columns int,r int,c int,mark int,group_id int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if grid[r][c] == 0 || grid[r][c] > 1{
		return
	}
	grid[r][c] = mark
	var g Group
	g.group_id = group_id
	g.square = mark
	groups[r][c] = g
	mark_largestIsland(grid,groups,rows,columns,r - 1,c,mark,group_id)
	mark_largestIsland(grid,groups,rows,columns,r + 1,c,mark,group_id)
	mark_largestIsland(grid,groups,rows,columns,r,c - 1,mark,group_id)
	mark_largestIsland(grid,groups,rows,columns,r,c + 1,mark,group_id)
}

func calc_around(groups [][]Group,rows int,columns int,r int,c int)int{
	var record map[int]int = make(map[int]int)
	if (r - 1) >= 0{
		record[groups[r - 1][c].group_id] = groups[r - 1][c].square
	}
	if (r + 1) < rows{
		if _,ok := record[groups[r + 1][c].group_id];!ok{
			record[groups[r + 1][c].group_id] = groups[r + 1][c].square
		}
	}
	if (c - 1) >= 0{
		if _,ok := record[groups[r][c - 1].group_id];!ok{
			record[groups[r][c - 1].group_id] = groups[r][c - 1].square
		}
	}
	if (c + 1) < columns{
		if _,ok := record[groups[r][c + 1].group_id];!ok{
			record[groups[r][c + 1].group_id] = groups[r][c + 1].square
		}
	}
	var around int = 0
	for _,square := range record{
		around += square
	}
	return around
}

type Group struct{
	group_id int
	square int
}

func LargestIsland(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var groups [][]Group = make([][]Group,rows)
	for i := 0;i < rows;i++{
		groups[i] = make([]Group,columns)
	}
	var group_id int = 1
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 1{
				square := calc_largestIsland(grid,rows,columns,i,j)
				mark_largestIsland(grid,groups,rows,columns,i,j,square,group_id)
				group_id++
			}
		}
	}
	var res int = 0
	var has_zero bool = false
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 0{
				has_zero = true
				cur := calc_around(groups,rows,columns,i,j) + 1
				if cur > res{
					res = cur
				}
			}
		}
	}
	if !has_zero{
		return rows * columns
	}
	return res
}