package array

import (
	"container/list"
	"math"
)

type point struct {
	x int
	y int
}

func UpdateMatrix(matrix [][]int) [][]int{
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var q list.List
	var res [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
		for j := 0;j < columns;j++{
			if matrix[i][j] == 0{
				var p point
				p.x = i
				p.y = j
				q.PushBack(p)
			}else{
				res[i][j] = math.MaxInt32
			}
		}
	}
	var dir [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var steps int = 1
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var p point = q.Front().Value.(point)
			q.Remove(q.Front())
			for _,d := range dir{
				var new_p point
				new_p.x = p.x + d[0]
				new_p.y = p.y + d[1]
				if new_p.x < 0 || new_p.x >= rows || new_p.y < 0 || new_p.y >= columns ||
					res[new_p.x][new_p.y] == 0 || res[new_p.x][new_p.y] < steps{
					continue
				}
				res[new_p.x][new_p.y] = min_int(res[new_p.x][new_p.y],steps)
				q.PushBack(new_p)
			}
		}
		steps++
	}
	return res
}

//func UpdateMatrix(matrix [][]int) [][]int {
//	var rows int = len(matrix)
//	var columns int = len(matrix[0])
//	var queue []point
//	for r := 0;r < rows;r++{
//		for c := 0;c < columns;c++{
//			if matrix[r][c] != 0{
//				matrix[r][c] = math.MaxInt32
//			}else{
//				var p point
//				p.x = r
//				p.y = c
//				queue = append(queue,p)
//			}
//		}
//	}
//	var dir [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
//	var distance int = 1
//	for len(queue) > 0{
//		var l int = len(queue)
//		for i := 0;i < l;i++{
//			var p point = queue[0]
//			queue = queue[1:]
//			for _,d := range dir{
//				r := p.x + d[0]
//				c := p.y + d[1]
//				if r < 0 || r >= rows || c < 0 || c >= columns{
//					continue
//				}
//				if matrix[r][c] < distance{
//					continue
//				}
//				matrix[r][c] = distance
//				var next_p point
//				next_p.x = r
//				next_p.y = c
//				queue = append(queue,next_p)
//			}
//		}
//		distance++
//	}
//	return matrix
//}