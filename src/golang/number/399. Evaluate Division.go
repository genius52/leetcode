package number

func dfs_calcEquation(start string,end string,pre_val float64,graph map[string][]string,express map[string]float64,visited map[string]bool)float64{
	if _,ok := graph[start];!ok{
		return -1
	}
	if _,ok := visited[start];ok{
		return -1
	}
	if start == end{
		return 1
	}
	visited[start] = true
	key := start + "," + end
	if val,ok := express[key];ok{
		return val * pre_val
	}
	for _,next := range graph[start]{
		k := start + "," + next
		if val,ok := express[k];ok{
			ret := dfs_calcEquation(next,end,val * pre_val,graph,express,visited)
			if ret != -1{
				return ret
			}
		}
	}
	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var graph map[string][]string = make(map[string][]string)
	var express map[string]float64 = make(map[string]float64)
	var l1 int = len(equations)
	for i := 0;i < l1;i++{
		graph[equations[i][0]] = append(graph[equations[i][0]],equations[i][1])
		graph[equations[i][1]] = append(graph[equations[i][1]],equations[i][0])
		key1 := equations[i][0] + "," + equations[i][1]
		express[key1] = values[i]
		key2 := equations[i][1] + "," + equations[i][0]
		express[key2] = 1/values[i]
	}
	var l2 int = len(queries)
	var res []float64 = make([]float64,l2)
	for i := 0;i < l2;i++{
		var visited map[string]bool = make(map[string]bool)
		ret := dfs_calcEquation(queries[i][0],queries[i][1],1,graph,express,visited)
		res[i] = ret
	}
	return res
}