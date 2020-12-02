package diagram

//Input: ["c==c","b==d","x!=z"]
//Output: true
//Input: ["a==b","b!=c","c==a"]
//Output: false
func dfs_equationsPossible(src uint8,dst uint8,graph map[uint8]map[uint8]bool,visited []bool)bool{
	if visited[src]{
		return true
	}
	visited[src] = true
	if _,ok := graph[src];ok{
		if _,ok2 := graph[src][dst];ok2{
			return true
		}
		for next,_ := range graph[src]{
			if visited[next]{
				continue
			}
			res := dfs_equationsPossible(next,dst,graph,visited)
			if res{
				return true
			}
		}
	}
	return false
}

func EquationsPossible(equations []string) bool {
	var graph map[uint8]map[uint8]bool = make(map[uint8]map[uint8]bool)
	var l int = len(equations)
	for i := 0;i < l;i++{
		if equations[i][1] == '=' {
			var first uint8 = equations[i][0] - 'a'
			var second uint8 = equations[i][3] - 'a'
			if _,ok := graph[first];!ok{
				graph[first] = make(map[uint8]bool)
			}
			graph[first][second] = true
			if _,ok := graph[second];!ok{
				graph[second] = make(map[uint8]bool)
			}
			graph[second][first] = true
		}
	}
	for i := 0;i < l;i++{
		if equations[i][1] == '!'{
			var first uint8 = equations[i][0] - 'a'
			var second uint8 = equations[i][3] - 'a'
			if first == second{
				return false
			}
			var visited []bool = make([]bool,26)
			equal := dfs_equationsPossible(first,second,graph,visited)
			if equal{
				return false
			}
		}
	}
	return true
}