package diagram

import (
	"container/list"
	"strconv"
)

type Status struct{
	x int
	y int
	key_cnt [6]bool
}

func ShortestPathAllKeys(grid []string) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var target_key_cnt [6]bool
	var q list.List
	var start_x int = 0
	var start_y int = 0
	var t []string = make([]string,rows * columns)
	var idx int =0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] >= 'a' && grid[i][j] <= 'f'{
				target_key_cnt[grid[i][j] - 'a'] = true
			}
			if grid[i][j] == '@'{
				start_x = i
				start_y = j
			}
			t[idx] = string(grid[i][j])
			idx++
		}
	}
	var visited map[string]bool = make(map[string]bool)
	var start Status
	start.x = start_x
	start.y = start_y
	visited["0 " + strconv.Itoa(start.x) + " " + strconv.Itoa(start.y)] = true
	//t[start_x * columns + start_y] = "1"
	//start.trace = strings.Join(t,"")
	q.PushBack(start)
	var dirs [4][2]int = [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur Status = q.Front().Value.(Status)
			q.Remove(q.Front())
			if cur.key_cnt[0] == target_key_cnt[0] && cur.key_cnt[1] == target_key_cnt[1] && cur.key_cnt[2] == target_key_cnt[2] &&
				cur.key_cnt[3] == target_key_cnt[3] && cur.key_cnt[4] == target_key_cnt[4] && cur.key_cnt[5] == target_key_cnt[5]{
				return steps
			}
			for _,dir := range dirs{
				var next Status
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				next.key_cnt = cur.key_cnt
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns{
					continue
				}
				if grid[next.x][next.y] == '#'{
					continue
				}
				if grid[next.x][next.y] >= 'A' && grid[next.x][next.y] <= 'F'{
					if !next.key_cnt[grid[next.x][next.y] - 'A']{
						continue
					}
				}
				if grid[next.x][next.y] >= 'a' && grid[next.x][next.y] <= 'z'{
					next.key_cnt[grid[next.x][next.y] - 'a'] = true
				}
				//keys from -> to
				c := 0
				for i := 0;i < 6;i++{
					if next.key_cnt[i]{
						c |= 1 << i
					}
				}
				key := strconv.Itoa(c) + " " + strconv.Itoa(next.x) + " " + strconv.Itoa(next.y);
				if _,ok := visited[key];ok{
					continue
				}
				visited[key] = true
				q.PushBack(next)
			}
		}
		steps++
	}
	return -1
}