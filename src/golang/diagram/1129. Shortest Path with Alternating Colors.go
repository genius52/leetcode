package diagram

import (
	"container/list"
)

//Input: n = 3, red_edges = [[0,1]], blue_edges = [[1,2]]
//Output: [0,1,2]
func bfs_shortestAlternatingPaths(n int,red_graph [][]bool,blue_graph [][]bool,is_red bool,res []int){
	steps := 1
	var q list.List
	q.PushBack(0)
	var visited [][][2]bool = make([][][2]bool,n)//prevent circle
	for i := 0;i < n;i++{
		visited[i] = make([][2]bool,n)
	}
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			if is_red{
				for j := 0;j < len(red_graph[cur]);j++{
					if red_graph[cur][j] && !visited[cur][j][0] {
						if res[j] == -1{
							res[j] = steps
						}else{
							res[j] = min_int(res[j],steps)
						}
						q.PushBack(j)
						visited[cur][j][0] = true
					}
				}
			}else{
				for j := 0;j < len(blue_graph[cur]);j++{
					if blue_graph[cur][j] && !visited[cur][j][1] {
						if res[j] == -1{
							res[j] = steps
						}else{
							res[j] = min_int(res[j],steps)
						}
						q.PushBack(j)
						visited[cur][j][1]  = true
					}
				}
			}
		}
		is_red = !is_red
		steps++
	}
}

func ShortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	var red_graph [][]bool = make([][]bool, n)
	var blue_graph [][]bool = make([][]bool, n)
	for i := 0; i < n; i++ {
		red_graph[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		blue_graph[i] = make([]bool, n)
	}
	for _, red := range red_edges {
		red_graph[red[0]][red[1]] = true
	}
	for _, blue := range blue_edges {
		blue_graph[blue[0]][blue[1]] = true
	}
	var res []int = make([]int,n)
	for i := 0;i < n;i++{
		res[i] = -1
	}
	res[0] = 0
	bfs_shortestAlternatingPaths(n,red_graph,blue_graph,true,res)
	bfs_shortestAlternatingPaths(n,red_graph,blue_graph,false,res)
	return res
}


//const (
//	FIRST int = 0
//	SECOND int = 1
//)
//
//func bfs_shortestAlternatingPaths(target int,n int,first_graph [][]bool,second_graph [][]bool)int{
//	var q list.List
//	q.PushFront(0)
//	var visited [][][2]bool = make([][][2]bool,n)
//	for i := 0;i < n;i++{
//		visited[i] = make([][2]bool,n)
//	}
//	//visited[0] = true
//	var steps int = 1
//	var use_first bool = true
//	for q.Len() > 0{
//		var l int = q.Len()
//		for i := 0;i < l;i++{
//			var start int = q.Back().Value.(int)
//			q.Remove(q.Back())
//			for j := 0;j < n;j++{
//				if use_first{
//					if visited[start][j][FIRST] {
//						continue
//					}
//				}else{
//					if visited[start][j][SECOND] {
//						continue
//					}
//				}
//
//				if use_first{
//					if first_graph[start][j]{
//						if j == target{
//							return steps
//						}
//						q.PushFront(j)
//						visited[start][j][FIRST] = true
//					}
//				}else{
//					if second_graph[start][j]{
//						if j == target{
//							return steps
//						}
//						q.PushFront(j)
//						visited[start][j][SECOND] = true
//					}
//				}
//			}
//		}
//		use_first = !use_first
//		steps++
//	}
//	return -1
//}
//
//func ShortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
//	var red_graph [][]bool = make([][]bool,n)
//	var blue_graph [][]bool = make([][]bool,n)
//	for i := 0;i < n;i++{
//		red_graph[i] = make([]bool,n)
//	}
//	for i := 0;i < n;i++{
//		blue_graph[i] = make([]bool,n)
//	}
//	for _,red := range red_edges{
//		red_graph[red[0]][red[1]] = true
//	}
//	for _,blue := range blue_edges{
//		blue_graph[blue[0]][blue[1]] = true
//	}
//	var res []int = make([]int,n)
//	res[0] = 0
//	for i := 1;i < n;i++{
//		red_first := bfs_shortestAlternatingPaths(i,n,red_graph,blue_graph)
//		blue_first := bfs_shortestAlternatingPaths(i,n,blue_graph,red_graph)
//		if red_first == -1 && blue_first == -1{
//			res[i] = -1
//		}else if red_first == -1{
//			res[i] = blue_first
//		}else if blue_first == -1{
//			res[i] = red_first
//		}else{
//			res[i] = int(math.Min(float64(red_first),float64(blue_first)))
//		}
//	}
//	return res
//}