package array

import "container/list"

//DFS solution
func dfs_containsCycle(grid *[][]byte,rows int,columns int,r int,c int,parent_r int,parent_c int,val byte,visited *[][]bool)bool{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return false
	}
	if (*visited)[r][c]{
		return false
	}
	(*visited)[r][c] = true
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _,dir := range dirs{
		next_r := r + dir[0]
		next_c := c + dir[1]
		if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns || (*grid)[next_r][next_c] != val{
			continue
		}
		if (*visited)[next_r][next_c] && next_r != parent_r && next_c != parent_c{
			return true
		}
		ret := dfs_containsCycle(grid,rows,columns,next_r,next_c,r,c,val,visited)
		if ret{
			return true
		}
	}
	return false
}

func containsCycle(grid [][]byte) bool{
	var rows int = len(grid)
	var columns int = len(grid[0])
	var visited [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		visited[i] = make([]bool,columns)
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if !visited[i][j]{
				ret := dfs_containsCycle(&grid,rows,columns,i,j,-1,-1,grid[i][j],&visited)
				if ret{
					return true
				}
			}
		}
	}
	return false
}

//BFS solution
type Path struct{
	val int
	r int
	c int
	parent_r int
	parent_c int
}

func ContainsCycle(grid [][]byte) bool {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var trace [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		trace[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++ {
		for j := 0; j < columns; j++ {
			trace[i][j] = int(grid[i][j])
		}
	}
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var visited_tag int = 1000
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if trace[i][j] >= 'a' && trace[i][j] <= 'z'{
				//var val byte = grid[i][j]
				var q list.List
				var first Path
				first.val = trace[i][j]
				first.r = i
				first.c = j
				first.parent_r = -1
				first.parent_c = -1
				q.PushBack(first)
				trace[i][j] = visited_tag
				for q.Len() > 0{
					var cur_len int = q.Len()
					for i := 0;i < cur_len;i++{
						var cur Path = q.Front().Value.(Path)
						q.Remove(q.Front())
						for _,dir := range dirs{
							var next_r int = cur.r + dir[0]
							var next_c int = cur.c + dir[1]
							if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns || (cur.parent_r == next_r && cur.parent_c == next_c) {
								continue
							}
							if trace[next_r][next_c] >= 'a' && trace[next_r][next_c] <= 'z' && trace[next_r][next_c] != cur.val{
								continue
							}
							if !(trace[next_r][next_c] >= 'a' && trace[next_r][next_c] <= 'z') && trace[next_r][next_c] < visited_tag{
								continue
							}
							if trace[next_r][next_c] == visited_tag && (next_r != cur.parent_r || next_c != cur.parent_c){
								return true
							}
							trace[next_r][next_c] = visited_tag
							var next_path Path
							next_path.val = cur.val
							next_path.r = next_r
							next_path.c = next_c
							next_path.parent_r = cur.r
							next_path.parent_c = cur.c
							q.PushBack(next_path)
						}
					}
				}
				visited_tag++
			}
		}
	}
	return false
}