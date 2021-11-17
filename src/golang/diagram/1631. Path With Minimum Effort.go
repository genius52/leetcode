package diagram

import (
	"container/list"
)

//bellman ford
type Pos_weight struct {
	r int
	c int
	weight int
}

func MinimumEffortPath(heights [][]int) int{
	var rows int = len(heights)
	var columns int = len(heights[0])
	var weights [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		weights[i] = make([]int,columns)
		for j := 0;j < columns;j++{
			weights[i][j] = 1 << 32 - 1
		}
	}
	weights[0][0] = 0
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var q list.List
	var first Pos_weight = Pos_weight{0,0,0}
	q.PushBack(first)
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			cur := q.Front().Value.(Pos_weight)
			q.Remove(q.Front())
			for _,dir := range dirs{
				var next Pos_weight
				next.r = cur.r + dir[0]
				next.c = cur.c + dir[1]
				if next.r < 0 || next.r >= rows || next.c < 0 || next.c >= columns{
					continue
				}
				next_weight := max_int(cur.weight,abs_int(heights[cur.r][cur.c] - heights[next.r][next.c]))
				if next_weight >= weights[next.r][next.c]{
					continue
				}
				weights[next.r][next.c] = next_weight
				next.weight = next_weight
				q.PushBack(next)
			}
		}
	}
	return weights[rows - 1][columns - 1]
}