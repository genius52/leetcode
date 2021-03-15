package diagram

func findCenter(edges [][]int) int {
	var l int = len(edges)
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	var n int = l + 1
	for i := 0;i < l;i++{
		if _,ok := graph[edges[i][0]];!ok{
			graph[edges[i][0]] = make(map[int]bool)
		}
		graph[edges[i][0]][edges[i][1]] = true
		if _,ok := graph[edges[i][1]];!ok{
			graph[edges[i][1]] = make(map[int]bool)
		}
		graph[edges[i][1]][edges[i][0]] = true
		if edges[i][0] > n{
			n = edges[i][0]
		}
		if edges[i][1] > n{
			n = edges[i][1]
		}
	}
	var res int = 0
	for k,v := range graph{
		if len(v) == l{
			res = k
			break
		}
	}
	return res
}