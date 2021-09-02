package diagram

import (
	"container/list"
	"strconv"
)

type Trace struct {
	x1 int
	y1 int
	x2 int
	y2 int
	vertical bool
}

func MinimumMoves(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var q list.List
	var visited map[string]bool = make(map[string]bool)
	visited["0,0,0,1"] = true
	var t Trace
	t.x1 = 0
	t.y1 = 0
	t.x2 = 0
	t.y2 = 1
	t.vertical = false
	q.PushBack(t)
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(Trace)
			//fmt.Println(strconv.Itoa(cur.x1) + "," + strconv.Itoa(cur.y1) + " | " + strconv.Itoa(cur.x2) + "," + strconv.Itoa(cur.y2))
			q.Remove(q.Front())
			if (cur.x1 == rows - 1 && cur.y1 == columns - 1 && cur.x2 == rows - 1 && cur.y2 == columns - 2) ||
				(cur.x1 == rows - 1 && cur.y1 == columns - 2 && cur.x2 == rows - 1 && cur.y2 == columns - 1){
				return steps
			}
			if cur.vertical {//vertical to horizon
				var next Trace
				next.x1 = min_int(cur.x1,cur.x2)
				next.y1 = cur.y1
				next.x2 = next.x1
				next.y2 = cur.y1 + 1
				next.vertical = false
				if next.x1 >= 0 && next.x1 < rows && next.x2 >= 0 && next.x2 < rows &&
					next.y1 >= 0 && next.y1 < columns && next.y2 >= 0 && next.y2 < columns &&
					grid[next.x1][next.y1] != 1 && grid[next.x2][next.y2] != 1 && grid[max_int(cur.x1,cur.x2)][next.y2] != 1{
					key := strconv.Itoa(next.x1) + "," + strconv.Itoa(next.y1) + "," +
						strconv.Itoa(next.x2) + "," + strconv.Itoa(next.y2)
					if _,ok := visited[key];!ok{
						visited[key] = true
						q.PushBack(next)
					}
				}
			}else{//horizon to vertical
				var next Trace
				next.x1 = cur.x1
				next.y1 = min_int(cur.y1,cur.y2)
				next.x2 = next.x1 + 1
				next.y2 = cur.y1
				next.vertical= true
				if next.x1 >= 0 && next.x1 < rows && next.x2 >= 0 && next.x2 < rows &&
					next.y1 >= 0 && next.y1 < columns && next.y2 >= 0 && next.y2 < columns &&
					grid[next.x1][next.y1] != 1 && grid[next.x2][next.y2] != 1 && grid[next.x2][max_int(cur.y1,cur.y2)] != 1{
					key := strconv.Itoa(next.x1) + "," + strconv.Itoa(next.y1) + "," +
						strconv.Itoa(next.x2) + "," + strconv.Itoa(next.y2)
					if _,ok := visited[key];!ok{
						visited[key] = true
						q.PushBack(next)
					}
				}
			}
			//move down
			var next1 Trace = cur
			next1.x1 = cur.x1 + 1
			next1.x2 = cur.x2 + 1
			if next1.x1 >= 0 && next1.x1 < rows && next1.x2 >= 0 && next1.x2 < rows &&
				next1.y1 >= 0 && next1.y1 < columns && next1.y2 >= 0 && next1.y2 < columns &&
				grid[next1.x1][next1.y1] != 1 && grid[next1.x2][next1.y2] != 1{
				key := strconv.Itoa(next1.x1) + "," + strconv.Itoa(next1.y1) + "," +
					strconv.Itoa(next1.x2) + "," + strconv.Itoa(next1.y2)
				if _,ok := visited[key];!ok{
					visited[key] = true
					q.PushBack(next1)
				}
			}
			//move right
			var next2 Trace = cur
			next2.y1 = cur.y1 + 1
			next2.y2 = cur.y2 + 1
			if next2.x1 >= 0 && next2.x1 < rows && next2.x2 >= 0 && next2.x2 < rows &&
				next2.y1 >= 0 && next2.y1 < columns && next2.y2 >= 0 && next2.y2 < columns &&
				grid[next2.x1][next2.y1] != 1 && grid[next2.x2][next2.y2] != 1{
				key := strconv.Itoa(next2.x1) + "," + strconv.Itoa(next2.y1) + "," +
					strconv.Itoa(next2.x2) + "," + strconv.Itoa(next2.y2)
				if _,ok := visited[key];!ok{
					visited[key] = true
					q.PushBack(next2)
				}
			}
		}
		steps++
	}
	return -1
}