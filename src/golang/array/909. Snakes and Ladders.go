package array

import "container/list"

func SnakesAndLadders(board [][]int) int {
	var rows int = len(board)
	var columns int = len(board[0])
	var data []int = make([]int,rows * columns)
	var idx int = 0
	var left_to_right bool = true
	for i := rows - 1;i >= 0;i--{
		for j := 0;j < columns;j++{
			if left_to_right{
				data[idx] = board[i][j]
			}else{
				data[idx] = board[i][columns - 1 - j]
			}
			idx++
		}
		left_to_right = !left_to_right
	}
	var end int = rows * columns - 1
	var visited []bool = make([]bool,rows * columns)
	var q list.List
	q.PushBack(0)
	visited[0] = true
	var steps int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur int = q.Front().Value.(int)
			q.Remove(q.Front())
			if cur >= end{
				return steps
			}
			for j := 1;j <= 6 && (cur + j) <= end;j++{
				next := cur + j
				if data[next] != -1{
					if !visited[data[next] - 1]{
						visited[data[next] - 1] = true
						q.PushBack(data[next] - 1)
					}
				}else{
					if visited[next]{
						continue
					}
					visited[next] = true
					q.PushBack(next)
				}
			}
		}
		steps++
	}
	return -1
}

//func dp_snakesAndLadders(data []int,end int,cur_num int,memo []int,visited []bool)int{
//	if cur_num >= end{
//		return 0
//	}
//	if memo[cur_num] != 0{
//		return memo[cur_num]
//	}
//	visited[cur_num] = true
//	var res int = 2147483647
//	var longest_step int = -1
//	for step := 1;step <= 6;step++{
//		next := cur_num + step
//		if next >= end{
//			return 1
//		}
//		if visited[next]{
//			continue
//		}
//		if data[next] != -1{
//			if next != cur_num{
//				visited[next] = true
//				steps := 1 + dp_snakesAndLadders(data,end,data[next],memo,visited)
//				if steps < res{
//					res = steps
//				}
//			}
//		}else{
//			longest_step = step
//		}
//	}
//	if longest_step != -1{
//		steps := 1 + dp_snakesAndLadders(data,end,cur_num + longest_step,memo,visited)
//		if steps < res{
//			res = steps
//		}
//	}
//	memo[cur_num] = res
//	return res
//}
//
//func SnakesAndLadders(board [][]int) int {
//	var rows int = len(board)
//	var columns int = len(board[0])
//	var memo []int = make([]int,rows * columns)
//	var data []int = make([]int,rows * columns)
//	var idx int = 0
//	for i := 0;i < rows;i++{
//		for j := 0;j < columns;j++{
//			if (i | 1) != i{
//				data[idx] = board[i][j]
//			}else{
//				data[idx] = board[i][columns - 1 - j]
//			}
//			idx++
//		}
//	}
//	var visited []bool = make([]bool,rows * columns)
//	return dp_snakesAndLadders(data,rows * columns,0,memo,visited)
//}