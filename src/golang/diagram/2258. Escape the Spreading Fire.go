package diagram

import "container/list"

func check_maximumMinutes(grid [][]int,rows int,columns int,limit int)bool{
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var data [][]int = make([][]int,rows)
	//var fire [][]bool = make([][]bool,rows)
	var q1 list.List
	for i := 0;i < rows;i++{
		data[i] = make([]int,columns)
		//fire[i] = make([]bool,columns)
		for j := 0;j < columns;j++{
			data[i][j] = grid[i][j]
			if grid[i][j] == 1{
				//fire[i][j] = true
				var p Point
				p.x = i
				p.y = j
				q1.PushBack(p)
			}
		}
	}
	var steps int = 0
	for q1.Len() > 0 && steps < limit{
		var cur_len int = q1.Len()
		for i := 0;i < cur_len;i++{
			var cur Point = q1.Front().Value.(Point)
			q1.Remove(q1.Front())
			for _,dir := range dirs{
				var next Point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns ||
					data[next.x][next.y] == 2 || data[next.x][next.y] == 1{
					continue
				}
				data[next.x][next.y] = 1
				q1.PushBack(next)
			}
		}
		steps++
	}
	if data[0][0] == 1{
		return false
	}
	var q2 list.List
	var start Point
	start.x = 0
	start.y = 0
	q2.PushBack(start)
	for q2.Len() > 0{
		var q2_len int = q2.Len()
		for i := 0;i < q2_len;i++{
			var cur Point = q2.Front().Value.(Point)
			q2.Remove(q2.Front())
			if data[cur.x][cur.y] == 1{
				continue
			}
			data[cur.x][cur.y] = 1
			if  cur.x == rows - 1 && cur.y == columns - 1{
				return true
			}
			for _,dir := range dirs{
				var next Point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns ||
					data[next.x][next.y] == 1 || data[next.x][next.y] == 2{
					continue
				}
				if  next.x == rows - 1 && next.y == columns - 1{
					return true
				}
				q2.PushBack(next)
			}
		}
		var q1_len int = q1.Len()
		for i := 0;i < q1_len;i++{
			var cur Point = q1.Front().Value.(Point)
			q1.Remove(q1.Front())
			for _,dir := range dirs{
				var next Point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns ||
					data[next.x][next.y] == 2 || data[next.x][next.y] == 1{
					continue
				}
				data[next.x][next.y] = 1
				q1.PushBack(next)
			}
		}
	}
	return false
}

func MaximumMinutes(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var dis_from_start [][]int = make([][]int,rows)
	var dis_from_fire [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dis_from_start[i] = make([]int,columns)
		dis_from_fire[i] = make([]int,columns)
	}

	var left int = 0
	var right int = 1e9
	var res int = -1
	for left <= right{
		mid := (left + right)/2
		ret := check_maximumMinutes(grid,rows,columns,mid)
		if ret{
			left = mid + 1
			res = mid
		}else{
			right = mid - 1
			//res = right
		}
	}
	return res
}

//func maximumMinutes(grid [][]int) int {
//	var rows int = len(grid)
//	var columns int = len(grid[0])
//	var dis_from_start [][]int = make([][]int,rows)
//	var dis_from_fire [][]int = make([][]int,rows)
//	for i := 0;i < rows;i++{
//		dis_from_start[i] = make([]int,columns)
//		dis_from_fire[i] = make([]int,columns)
//	}
//	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
//	var q1 list.List
//	var start Point
//	start.x = 0
//	start.y = 0
//	q1.PushBack(start)
//	var steps1 int = 1
//	for q1.Len() > 0{
//		var cur_len int = q1.Len()
//		for i := 0;i < cur_len;i++{
//			var cur Point = q1.Front().Value.(Point)
//			q1.Remove(q1.Front())
//			for _,dir := range dirs{
//				var next Point
//				next.x = cur.x + dir[0]
//				next.y = cur.y + dir[1]
//				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns || grid[next.x][next.y] == 1 ||
//					grid[next.x][next.y] == 2 || dis_from_start[next.x][next.y] <= steps1{
//					continue
//				}
//				dis_from_start[next.x][next.y] = steps1
//				q1.PushBack(next)
//			}
//		}
//		steps1++
//	}
//	//bfs_maximumMinutes(grid,rows,columns,rows - 1,columns - 1,dis_from_fire)
//	var q2 list.List
//	var steps2 int = 1
//	for i := 0;i < rows;i++{
//		for j := 0;j < columns;j++{
//			if grid[i][j] == 2{
//				var p Point
//				p.x = i
//				p.y = j
//				q2.PushBack(p)
//			}
//		}
//	}
//	for q2.Len() > 0{
//		var cur_len int = q2.Len()
//		for i := 0;i < cur_len;i++{
//			var cur Point = q2.Front().Value.(Point)
//			q2.Remove(q2.Front())
//			for _,dir := range dirs{
//				var next Point
//				next.x = cur.x + dir[0]
//				next.y = cur.y + dir[1]
//				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns || grid[next.x][next.y] == 1 ||
//					grid[next.x][next.y] == 2 || dis_from_fire[next.x][next.y] <= steps2{
//					continue
//				}
//				dis_from_fire[next.x][next.y] = steps2
//				q2.PushBack(next)
//			}
//		}
//		steps2++
//	}
//	var diff int = dis_from_fire[rows - 1][columns - 1] - dis_from_start[rows - 1][columns - 1]
//	if diff <= 0{
//		return -1
//	}
//	var start_diff int = diff
//	var q3 list.List
//	var end Point
//	end.x = rows - 1
//	end.y = columns - 1
//	q3.PushBack(end)
//	for q3.Len() > 0{
//		var cur_len int = q3.Len()
//		for i := 0;i < cur_len;i++{
//			var cur Point = q3.Front().Value.(Point)
//			if cur.x == 0 && cur.y == 0{
//				return start_diff
//			}
//			q3.Remove(q3.Front())
//			for _,dir := range dirs{
//				var next Point
//				next.x = cur.x + dir[0]
//				next.y = cur.y + dir[1]
//				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns || grid[next.x][next.y] == 2{
//					continue
//				}
//				cur_diff := dis_from_fire[next.x][next.y] - dis_from_start[next.x][next.y]
//				if cur_diff > diff{
//
//				}
//			}
//		}
//		diff++
//	}
//	return -1
//}