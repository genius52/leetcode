package diagram

import (
	"container/list"
)

func OrangesRotting(grid [][]int) int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	var fresh int = 0
	var q list.List
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 1{
				fresh++
			}
			if grid[i][j] == 2{
				var p point
				p.x = i
				p.y = j
				q.PushBack(p)
			}
		}
	}
	if fresh == 0{
		return 0
	}
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur point = q.Front().Value.(point)
			q.Remove(q.Front())
			for _,d := range dirs{
				var next point
				next.x = cur.x + d[0]
				next.y = cur.y + d[1]
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns || grid[next.x][next.y] != 1{
					continue
				}
				grid[next.x][next.y] = 2
				q.PushBack(next)
				fresh--
			}
		}
		steps++
		if fresh == 0{
			return steps
		}
	}
	if fresh != 0{
		return -1
	}
	return steps
}

//func bfs2(grid [][]int,i int,j int,depth int) int{
//	var rows int = len(grid)
//	var columns int = len(grid[0])
//	if i < 0 || i >= rows || j < 0 || j >= columns || grid[i][j] != 1{
//		return 0
//	}
//	grid[i][j] = 2
//	depth++
//	max_depth := depth
//	res :=  bfs2(grid,i-1,j,depth)
//	if res > max_depth{
//		max_depth = res
//	}
//	res = bfs2(grid,i+1,j,depth)
//	if res > max_depth{
//		max_depth = res
//	}
//	res = bfs2(grid,i,j-1,depth)
//	if res > max_depth{
//		max_depth = res
//	}
//	res = bfs2(grid,i,j+1,depth)
//	if res > max_depth{
//		max_depth = res
//	}
//	return max_depth
//}
////[[2],[1],[1],[1],[2],[1[1]]
//func orangesRotting2(grid [][]int) int {
//	var rows int = len(grid)
//	var columns int = len(grid[0])
//	max_res := 0
//	for i := 0;i < rows;i++{
//		for j :=0;j < columns;j++{
//			if grid[i][j] == 2{
//				res := bfs2(grid,i-1,j,0)
//				if res > max_res{
//					max_res = res
//				}
//				res = bfs2(grid,i+1,j,0)
//				if res > max_res{
//					max_res = res
//				}
//				res = bfs2(grid,i,j-1,0)
//				if res > max_res{
//					max_res = res
//				}
//				res = bfs2(grid,i,j+1,0)
//				if res > max_res{
//					max_res = res
//				}
//			}
//		}
//	}
//	for i := 0;i < rows;i++{
//		for j :=0;j < columns;j++{
//			if grid[i][j] == 1{
//				return -1
//			}
//		}
//	}
//	return max_res
//}


//994
//Input: [[2,1,1],[1,1,0],[0,1,1]]
//Output: 4
//func fill_near_orgin(grid [][]int,visited,r,c,rows,cols int){
//	if r - 1 >= 0 && grid[r-1][c] == 1{
//		grid[r-1][c] = 2
//	}
//}
//
//func orangesRotting(grid [][]int) int {
//	rows := len(grid)
//	if rows == 0 {
//		return 0
//	}
//	columns := len(grid[0])
//	var steps int = 0
//	var unvisited map[string]bool = make(map[string]bool) //store unvisit good orange position
//	var depth [][]int //store current step bad orange
//	for row := 0;row < rows;row++{
//		for col := 0;col < columns;col++{
//			if grid[row][col] == 1{
//				k := strconv.Itoa(row)+","+strconv.Itoa(col)
//				unvisited[k] = true
//			}
//			if grid[row][col] == 2{
//				depth = append(depth,[]int{row,col})
//			}
//		}
//	}
//	if len(unvisited) == 0{// If no good orange exist,return 0
//		return 0
//	}
//	for len(depth) > 0{
//		l := len(depth)
//		var cur_depth [][]int = make([][]int,len(depth))
//		copy(cur_depth,depth)
//		depth = make([][]int,0)//clear last level bad orange
//		for i := 0;i < l;i++{
//			r := cur_depth[i][0]
//			c := cur_depth[i][1]
//			//check 4 direction's orange
//			if r - 1 >= 0 && grid[r-1][c] == 1{
//				grid[r-1][c] = 2
//				k := strconv.Itoa(r-1)+","+strconv.Itoa(c)
//				delete(unvisited,k)
//				depth = append(depth, []int{r-1,c})
//			}
//			if r + 1 < rows && grid[r+1][c] == 1{
//				grid[r+1][c] = 2
//				k := strconv.Itoa(r+1)+","+strconv.Itoa(c)
//				delete(unvisited,k)
//				depth = append(depth, []int{r+1,c})
//			}
//			if c - 1 >= 0 && grid[r][c-1] == 1{
//				grid[r][c-1] = 2
//				k := strconv.Itoa(r)+","+strconv.Itoa(c-1)
//				delete(unvisited,k)
//				depth = append(depth, []int{r,c-1})
//			}
//			if c + 1 < columns && grid[r][c+1] == 1{
//				grid[r][c+1] = 2
//				k := strconv.Itoa(r)+","+strconv.Itoa(c+1)
//				delete(unvisited,k)
//				depth = append(depth, []int{r,c+1})
//			}
//		}
//		steps++
//		if len(unvisited) == 0{// when there is no good orange,return steps
//			return steps
//		}
//	}
//	if len(unvisited) > 0{
//		return -1
//	}
//	return steps
//}