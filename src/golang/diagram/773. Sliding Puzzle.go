package diagram

import "container/list"

func check_valid(board [2][3]int)bool{
	return board[0][0] == 1 && board[0][1] == 2 && board[0][2] == 3 &&
			board[1][0] == 4 && board[1][1] == 5 && board[1][2] == 0
}

func SlidingPuzzle(board [][]int) int {
	var visited map[int]bool = make(map[int]bool)
	var q list.List
	var dirs [4][2]int = [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}
	var start [2][3]int
	start[0][0] = board[0][0]
	start[0][1] = board[0][1]
	start[0][2] = board[0][2]
	start[1][0] = board[1][0]
	start[1][1] = board[1][1]
	start[1][2] = board[1][2]
	q.PushBack(start)
	key := start[0][0] + 10 * start[0][1] + 100 * start[0][2] +
		1000 * start[1][0] + 10000 * start[1][1] + 100000 * start[1][2]
	visited[key] = true
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var node [2][3]int = q.Front().Value.([2][3]int)
			q.Remove(q.Front())
			if check_valid(node){
				return steps
			}
			var r int = 0
			var c int = 0
			for i := 0;i < 2;i++{
				var find bool = false
				for j := 0;j < 3;j++{
					if node[i][j] == 0{
						r = i
						c = j
						find = true
						break
					}
				}
				if find{
					break
				}
			}
			for i := 0;i < 4;i++{
				next_r := r + dirs[i][0]
				next_c := c + dirs[i][1]
				if next_r < 0 || next_r >= 2 || next_c < 0 || next_c >= 3{
					continue
				}
				var nextnode [2][3]int = node
				nextnode[next_r][next_c],nextnode[r][c] = nextnode[r][c],nextnode[next_r][next_c]
				key := nextnode[0][0] + 10 * nextnode[0][1] + 100 * nextnode[0][2] +
					1000 * nextnode[1][0] + 10000 * nextnode[1][1] + 100000 * nextnode[1][2]
				if _,ok := visited[key];!ok{
					visited[key] = true
					q.PushBack(nextnode)
				}
			}
		}
		steps++
	}
	return -1
}