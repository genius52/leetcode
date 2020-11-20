package diagram

import (
	"container/list"
)

//In a given 2D binary array A, there are two islands.  (An island is a 4-directionally connected group of 1s not connected to any other 1s.)
//
//Now, we may change 0s to 1s so as to connect the two islands together to form 1 island.
//
//Return the smallest number of 0s that must be flipped.  (It is guaranteed that the answer is at least 1.)
//Input: A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//Output: 1
type point struct {
	x int
	y int
}

func dfs_mark(A [][]int,rows int,columns int,r int,c int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if A[r][c] == 0{
		A[r][c] = 3
		return
	}
	if A[r][c] == 2 || A[r][c] == 3{
		return
	}
	A[r][c] = 2
	dfs_mark(A,rows,columns,r - 1,c)
	dfs_mark(A,rows,columns,r + 1,c)
	dfs_mark(A,rows,columns,r,c - 1)
	dfs_mark(A,rows,columns,r,c + 1)
}

func mark_point(A [][]int,val int,rows int,columns int,r int,c int,q *list.List)bool{
	var p point
	if r - 1 >= 0{
		if A[r - 1][c] == 1{
			return true
		}else if A[r - 1][c] == 0{
			A[r - 1][c] = val
			p.x = r - 1
			p.y = c
			q.PushBack(p)
		}
	}
	if r + 1 < rows{
		if A[r + 1][c] == 1{
			return true
		}else if A[r + 1][c] == 0{
			A[r + 1][c] = val
			p.x = r + 1
			p.y = c
			q.PushBack(p)
		}
	}
	if c - 1 >= 0{
		if A[r][c - 1] == 1{
			return true
		}else if A[r][c - 1] == 0{
			A[r][c - 1] = val
			p.x = r
			p.y = c - 1
			q.PushBack(p)
		}
	}
	if c + 1 < columns{
		if A[r][c + 1] == 1{
			return true
		}else if A[r][c + 1] == 0{
			A[r][c + 1] = val
			p.x = r
			p.y = c + 1
			q.PushBack(p)
		}
	}
	return false
}

func ShortestBridge(A [][]int) int {
	var rows int = len(A)
	var columns int = len(A[0])
	for i := 0;i < rows;i++{
		var find bool = false
		for j := 0;j < columns;j++{
			if A[i][j] == 0{
				continue
			}
			find = true
			dfs_mark(A,rows,columns,i,j)
			break
		}
		if find{
			break
		}
	}
	var q list.List
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if A[i][j] == 3{
				var p point
				p.x = i
				p.y = j
				q.PushBack(p)
			}
		}
	}
	var mark int = 4
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			p := q.Front().Value.(point)
			q.Remove(q.Front())
			reach := mark_point(A,mark,rows,columns,p.x,p.y,&q)
			if reach{
				return mark - 3
			}
		}
		mark++
	}
	return mark
}