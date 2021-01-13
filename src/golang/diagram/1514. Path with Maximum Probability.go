package diagram

import "container/list"

//Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
//Output: 0.25000
//Explanation: There are two paths from start to end, one having a probability of success = 0.2 and the other has 0.5 * 0.5 = 0.25.
type nodestate struct {
	num int
	prob float64
}
func MaxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	var l int = len(edges)
	var graph map[int]map[int]float64 = make(map[int]map[int]float64)
	for i := 0;i < l;i++{
		if _,ok := graph[edges[i][0]];!ok{
			graph[edges[i][0]] = make(map[int]float64)
		}
		graph[edges[i][0]][edges[i][1]] = succProb[i]
		if _,ok := graph[edges[i][1]];!ok{
			graph[edges[i][1]] = make(map[int]float64)
		}
		graph[edges[i][1]][edges[i][0]] = succProb[i]
	}
	var fromstart []float64 = make([]float64,n)
	fromstart[start] = 1.0
	//for i := 0;i < n;i++{
	//	fromstart[i] =
	//}
	if _,ok := graph[start];ok{
		if _,ok2 := graph[start][end];ok2{
			fromstart[end] = graph[start][end]
		}
	}
	var q list.List
	var cur nodestate
	cur.num = start
	cur.prob = 1.0
	q.PushBack(cur)
	for q.Len() > 0{
		pre_node := q.Front().Value.(nodestate)
		q.Remove(q.Front())
		for neighbour,prob := range graph[pre_node.num]{
			if fromstart[neighbour] == 0{
				fromstart[neighbour] = pre_node.prob * prob
			}else{
				if pre_node.prob * prob > fromstart[neighbour]{
					fromstart[neighbour] = pre_node.prob * prob
				}else{
					continue
				}
			}
			var cur_node nodestate
			cur_node.num = neighbour
			cur_node.prob = fromstart[neighbour]
			q.PushBack(cur_node)
		}
	}
	return fromstart[end]
}