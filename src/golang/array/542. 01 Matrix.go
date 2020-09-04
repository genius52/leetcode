package array

import "math"

type point struct {
	x int
	y int
}

func UpdateMatrix(matrix [][]int) [][]int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var queue []point
	for r := 0;r < rows;r++{
		for c := 0;c < columns;c++{
			if matrix[r][c] != 0{
				matrix[r][c] = math.MaxInt32
			}else{
				var p point
				p.x = r
				p.y = c
				queue = append(queue,p)
			}
		}
	}
	var dir [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var distance int = 1
	for len(queue) > 0{
		var l int = len(queue)
		for i := 0;i < l;i++{
			var p point = queue[0]
			queue = queue[1:]
			for _,d := range dir{
				r := p.x + d[0]
				c := p.y + d[1]
				if r < 0 || r >= rows || c < 0 || c >= columns{
					continue
				}
				if matrix[r][c] < distance{
					continue
				}
				matrix[r][c] = distance
				var next_p point
				next_p.x = r
				next_p.y = c
				queue = append(queue,next_p)
			}
		}
		distance++
	}
	return matrix
}