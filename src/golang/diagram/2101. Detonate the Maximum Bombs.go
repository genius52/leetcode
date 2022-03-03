package diagram

//func check_maximumDetonation(bombs [][]int,i int,j int)bool{
//	delta_x := bombs[i][0] - bombs[j][0]
//	delta_y := bombs[i][1] - bombs[j][1]
//	return  (delta_x * delta_x + delta_y * delta_y) <= (bombs[i][2] * bombs[i][2])
//}

func dfs_maximumDetonation(graph [][]int,cur int,visited []bool)int{
	if visited[cur]{
		return 0
	}
	visited[cur] = true
	var res int = 1
	for i := 0;i < len(graph[cur]);i++{
		if visited[graph[cur][i]]{
			continue
		}
		res += dfs_maximumDetonation(graph,graph[cur][i],visited)
	}
	return res
}
func MaximumDetonation(bombs [][]int) int {
	var l int = len(bombs)
	var graph [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		for j := 0;j < l;j++{
			if i == j{
				continue
			}
			delta_x := bombs[i][0] - bombs[j][0]
			delta_y := bombs[i][1] - bombs[j][1]
			if (bombs[i][2] * bombs[i][2]) >= (delta_x * delta_x + delta_y * delta_y) {
				graph[i] = append(graph[i],j)
			}
			//if check_maximumDetonation(bombs,i,j){
			//	graph[i] = append(graph[i],j)
			//}
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		var visited []bool = make([]bool,l)
		ret := dfs_maximumDetonation(graph,i,visited)
		if ret > res{
			res = ret
		}
	}
	return res
}